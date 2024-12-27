package vehicle_request_poller

import (
	"context"
	"log"
	"sync"
	tripDomain "transportation/internal/trip/domain"
	tripPort "transportation/internal/trip/port"
	"transportation/internal/vehicle_request_poller/domain"

	"transportation/internal/vehicle_request_poller/port"
)

type Service struct {
	queue    port.QueueRepo
	tripRepo tripPort.Repo
}

func NewService(queue port.QueueRepo, tripRepo tripPort.Repo) port.Service {
	return &Service{queue: queue, tripRepo: tripRepo}
}

func (s *Service) PollAndPublish(ctx context.Context, req domain.PollerRequest) error {
	totalBatches := (req.TotalRecords + req.BatchSize - 1) / req.BatchSize
	var wg sync.WaitGroup
	sem := make(chan struct{}, req.ConcurrentJobs)

	j := 0
	for i := 0; i < totalBatches; i++ {
		offset := j * req.BatchSize
		j++
		if j == (req.ConcurrentJobs) {
			j = 0
		}

		sem <- struct{}{}

		wg.Add(1)

		go func(offset int) {
			defer wg.Done()
			defer func() {
				<-sem
			}()

			requests, err := s.tripRepo.GetShouldCheckVehicleRequests(ctx, req.BatchSize, offset)

			if err != nil {
				log.Printf("Error fetching ShouldCheckVehicleRequests: %v", err)
				return
			}

			requestIds := []tripDomain.VehicleRequestId{}

			for _, req := range requests {
				if err := s.queue.PublishRequest(ctx, req); err != nil {
					log.Printf("Failed to publish request ID %d: %v", req.ID, err)
				}
				requestIds = append(requestIds, req.ID)
			}

			err = s.tripRepo.UpdateVehicleRequestsLastCheckTime(ctx, requestIds)
			if err != nil {
				log.Printf("Error UpdateVehicleRequestsLastCheckTime: %v", err)
				return
			}

		}(offset)
	}

	wg.Wait()
	return nil
}
