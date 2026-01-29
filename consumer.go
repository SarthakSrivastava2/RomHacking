package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Item represents a work item in our system
type Item struct {
	ID    int
	Value int
}

// ProcessedItem represents a processed work item
type ProcessedItem struct {
	OriginalItem Item
	Result       int
	ProcessedBy  int
}

// SafeQueue is a thread-safe queue implementation using channels
type SafeQueue struct {
	items chan interface{}
}

// NewSafeQueue creates a new thread-safe queue with specified capacity
func NewSafeQueue(capacity int) *SafeQueue {
	return &SafeQueue{
		items: make(chan interface{}, capacity),
	}
}

// Enqueue adds an item to the queue
func (q *SafeQueue) Enqueue(item interface{}) {
	q.items <- item
}

// Dequeue removes and returns an item from the queue
func (q *SafeQueue) Dequeue() interface{} {
	return <-q.items
}

// Close closes the queue
func (q *SafeQueue) Close() {
	close(q.items)
}

// Producer generates items and puts them in the input queue
func Producer(id int, inputQueue *SafeQueue, numItems int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for i := 0; i < numItems; i++ {
		item := Item{
			ID:    id*1000 + i,
			Value: rand.Intn(100),
		}
		fmt.Printf("Producer %d produced item %d with value %d\n", id, item.ID, item.Value)
		inputQueue.Enqueue(item)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
	}
	fmt.Printf("Producer %d finished\n", id)
}

// Consumer processes items from the input queue and puts results in the output queue
func Consumer(id int, inputQueue, outputQueue *SafeQueue, done chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for {
		select {
		case <-done:
			fmt.Printf("Consumer %d received shutdown signal\n", id)
			return
		default:
			// Try to get an item from the input queue with a timeout
			select {
			case itemInterface := <-inputQueue.items:
				item := itemInterface.(Item)
				
				// Process the item (in this case, just square the value)
				result := item.Value * item.Value
				
				processedItem := ProcessedItem{
					OriginalItem: item,
					Result:       result,
					ProcessedBy:  id,
				}
				
				fmt.Printf("Consumer %d processed item %d: %d -> %d\n", 
					id, item.ID, item.Value, result)
				
				outputQueue.Enqueue(processedItem)
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(200)))
			case <-time.After(10 * time.Millisecond):
				// Just a small timeout to prevent busy waiting
				continue
			}
		}
	}
}

// ResultCollector collects and displays results from the output queue
func ResultCollector(outputQueue *SafeQueue, done chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	
	count := 0
	for {
		select {
		case <-done:
			fmt.Printf("Result collector received shutdown signal after processing %d items\n", count)
			return
		default:
			// Try to get a processed item from the output queue with a timeout
			select {
			case itemInterface := <-outputQueue.items:
				processedItem := itemInterface.(ProcessedItem)
				fmt.Printf("Result: Item %d with value %d was processed by Consumer %d, result: %d\n",
					processedItem.OriginalItem.ID, 
					processedItem.OriginalItem.Value,
					processedItem.ProcessedBy,
					processedItem.Result)
				count++
			case <-time.After(10 * time.Millisecond):
				// Just a small timeout to prevent busy waiting
				continue
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	// Configuration
	numProducers := 3
	numConsumers := 2
	itemsPerProducer := 5
	queueCapacity := 10
	
	// Create queues
	inputQueue := NewSafeQueue(queueCapacity)
	outputQueue := NewSafeQueue(queueCapacity)
	
	// Wait groups for synchronization
	var producerWg sync.WaitGroup
	var consumerWg sync.WaitGroup
	var collectorWg sync.WaitGroup
	
	// Channel for signaling consumers and collector to stop
	done := make(chan bool)
	
	// Start producers
	producerWg.Add(numProducers)
	for i := 0; i < numProducers; i++ {
		go Producer(i+1, inputQueue, itemsPerProducer, &producerWg)
	}
	
	// Start consumers
	consumerWg.Add(numConsumers)
	for i := 0; i < numConsumers; i++ {
		go Consumer(i+1, inputQueue, outputQueue, done, &consumerWg)
	}
	
	// Start result collector
	collectorWg.Add(1)
	go ResultCollector(outputQueue, done, &collectorWg)
	
	// Wait for all producers to finish
	producerWg.Wait()
	fmt.Println("All producers have finished")
	
	// Give consumers some time to process remaining items
	time.Sleep(1 * time.Second)
	
	// Signal consumers and collector to stop
	close(done)
	
	// Wait for consumers and collector to finish
	consumerWg.Wait()
	collectorWg.Wait()
	
	fmt.Println("All done!")
}
