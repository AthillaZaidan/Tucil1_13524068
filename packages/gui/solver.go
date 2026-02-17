package gui

import (
	bruteforce "Tucil1/packages/bruteforce"
	bruteforceoptimized "Tucil1/packages/bruteforce-optimized"
	"fmt"
	"time"
)

// Interval minimum antara GUI refresh (biar ga freeze)
const guiRefreshInterval = 30 * time.Millisecond

// runSolver — jalankan solver di goroutine terpisah (mode 1 = pure, mode 2 = optimized)
func (qa *QueensApp) runSolver(mode int) {
	qa.mu.Lock()
	qa.solving = true
	qa.stopFlag = false
	qa.lastSolution = nil
	qa.latestPlacement = nil
	qa.latestIter = 0
	qa.mu.Unlock()

	qa.clearQueens()
	qa.statusLabel.SetText("Status: Solving...")
	qa.iterLabel.SetText("Iterations: 0")
	qa.timeLabel.SetText("Time: 0 ms")

	// reset grid ke original
	for i := 0; i < qa.row; i++ {
		for j := 0; j < qa.col; j++ {
			qa.grid[i][j] = qa.originalGrid[i][j]
		}
	}

	startTime := time.Now()

	// =============================================
	// CHANNEL: solver kirim data → GUI ticker baca
	// =============================================
	// Solver CUMA nulis ke shared state (cepet, no blocking).
	// GUI ticker jalan terpisah, baca shared state tiap interval, update visual.
	// Ini yang bikin solver ga ke-slow-down sama GUI refresh.

	// start GUI updater ticker — jalan tiap guiRefreshInterval
	done := make(chan bool, 1)
	go qa.guiUpdaterLoop(startTime, done)

	var solution []int
	var found bool

	// bikin lightweight callback — cuma update shared state, no GUI call
	lightCallback := func(placement []int, iter int) {
		qa.mu.Lock()
		// copy placement biar ga race condition
		cpy := make([]int, len(placement))
		copy(cpy, placement)
		qa.latestPlacement = cpy
		qa.latestIter = iter
		qa.mu.Unlock()
	}

	if mode == 1 {
		bruteforce.OnStep = lightCallback
		bruteforce.StopFlag = &qa.stopFlag
		solution, found = bruteforce.Bruteforce_solve(qa.grid, qa.row, qa.col)
		bruteforce.OnStep = nil
		bruteforce.StopFlag = nil
	} else {
		bruteforceoptimized.OnStep = lightCallback
		bruteforceoptimized.StopFlag = &qa.stopFlag
		solution, found = bruteforceoptimized.Bruteforce_optimized_solve(qa.grid, qa.row, qa.col)
		bruteforceoptimized.OnStep = nil
		bruteforceoptimized.StopFlag = nil
	}

	// stop GUI updater
	done <- true

	duration := time.Since(startTime)

	// final update
	if found {
		qa.updateGridQueens(solution, len(solution))
		qa.statusLabel.SetText("Status: Solution Found!")
		qa.lastSolution = solution
	} else {
		qa.mu.Lock()
		stopped := qa.stopFlag
		qa.mu.Unlock()

		if stopped {
			qa.statusLabel.SetText("Status: Stopped")
		} else {
			qa.statusLabel.SetText("Status: No Solution Found")
		}
	}

	qa.timeLabel.SetText(fmt.Sprintf("Time: %d ms", duration.Milliseconds()))

	qa.mu.Lock()
	qa.solving = false
	qa.mu.Unlock()
}

// guiUpdaterLoop — goroutine terpisah yang update GUI pada interval tetap
// Solver thread TIDAK pernah sentuh GUI langsung — cuma update angka di shared state.
// Loop ini yang baca shared state dan refresh visual.
func (qa *QueensApp) guiUpdaterLoop(startTime time.Time, done chan bool) {
	ticker := time.NewTicker(guiRefreshInterval)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			// baca latest state dari solver
			qa.mu.Lock()
			placement := qa.latestPlacement
			iter := qa.latestIter
			stop := qa.stopFlag
			delay := qa.stepDelay
			qa.mu.Unlock()

			if stop {
				continue
			}

			// update visual
			if placement != nil {
				qa.updateGridQueens(placement, len(placement))
			}

			elapsed := time.Since(startTime)
			qa.iterLabel.SetText(fmt.Sprintf("Iterations: %d", iter))
			qa.timeLabel.SetText(fmt.Sprintf("Time: %d ms", elapsed.Milliseconds()))

			// kalo user mau slow-down buat liat step by step, tambahin delay
			if delay > guiRefreshInterval {
				time.Sleep(delay - guiRefreshInterval)
			}
		}
	}
}
