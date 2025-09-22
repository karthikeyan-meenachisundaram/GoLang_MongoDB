//Connect MongoDB with Go, Insert from Go, read from Mongo
//1. db.go
package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoURI returns the URI string (hidden from main.go)
func MongoURI() string {
	// You can also read this from an env variable here if you want
	return "mongodb+srv://Karthikmongodb.net/"
}

// Connect returns Mongo client, context, and cancel func.
// Caller must defer cancel() and client.Disconnect(ctx).
func Connect(timeoutSeconds int) (*mongo.Client, context.Context, context.CancelFunc, error) {
	uri := MongoURI() // get URI from this file

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSeconds)*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		cancel()
		return nil, nil, nil, err
	}

	// optional: ping
	if err := client.Ping(ctx, nil); err != nil {
		_ = client.Disconnect(ctx)
		cancel()
		return nil, nil, nil, err
	}

	fmt.Println("✅ Connected to MongoDB")
	return client, ctx, cancel, nil
}


//2. employee.go
package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Employee struct (same as before)
type Employee struct {
	EmpID  int     `bson:"empid"`
	Name   string  `bson:"name"`
	Salary float32 `bson:"salary"`
}

// FetchEmployees reads all docs using a fresh context for the operation.
func FetchEmployees(coll *mongo.Collection) error {
	// per-operation timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		return fmt.Errorf("find error: %w", err)
	}
	defer cursor.Close(ctx)

	var employees []Employee
	if err := cursor.All(ctx, &employees); err != nil {
		return fmt.Errorf("decode error: %w", err)
	}

	if len(employees) == 0 {
		fmt.Println("No employees found.")
		return nil
	}

	for _, e := range employees {
		fmt.Printf("EmpID: %d, Name: %s, Salary: %.2f\n", e.EmpID, e.Name, e.Salary)
	}
	return nil
}

// InsertEmployee inserts one Employee using a fresh context for the operation.
func InsertEmployee(coll *mongo.Collection, emp Employee) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := coll.InsertOne(ctx, emp)
	if err != nil {
		return fmt.Errorf("insert error: %w", err)
	}
	fmt.Println("Inserted Employee with ID:", res.InsertedID)
	return nil
}


//3. main.go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	client, ctx, cancel, err := Connect(10) // just call Connect, no URI here
	if err != nil {
		log.Fatal("Connection error:", err)
	}
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Println("disconnect error:", err)
		}
	}()

	coll := client.Database("my_db").Collection("Employee_Details")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Choose action: 1) Fetch  2) Insert")
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	choice, err := strconv.Atoi(line)
	if err != nil {
		log.Fatal("Invalid choice:", err)
	}

	switch choice {
	case 1:
		if err := FetchEmployees(coll); err != nil {
			log.Fatal("Fetch error:", err)
		}
	case 2:
		for {
			fmt.Print("Enter EmpID (integer): ")
			line, _ := reader.ReadString('\n')
			line = strings.TrimSpace(line)
			eid, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println("Invalid EmpID. Please enter an integer.")
				continue
			}

			fmt.Print("Enter Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			if name == "" {
				fmt.Println("Name cannot be empty.")
				continue
			}

			fmt.Print("Enter Salary (number): ")
			line, _ = reader.ReadString('\n')
			line = strings.TrimSpace(line)
			sal, err := strconv.ParseFloat(line, 32)
			if err != nil {
				fmt.Println("Invalid salary. Please enter a number.")
				continue
			}

			emp := Employee{EmpID: eid, Name: name, Salary: float32(sal)}
			if err := InsertEmployee(coll, emp); err != nil {
				log.Println("Insert error:", err)
			} else {
				fmt.Println("Inserted successfully.")
			}

			fmt.Print("Insert another? (y/n): ")
			yn, _ := reader.ReadString('\n')
			yn = strings.TrimSpace(strings.ToLower(yn))
			if yn != "y" && yn != "yes" {
				fmt.Println("Stopping inserts.")
				break
			}
		}
	default:
		fmt.Println("Invalid choice")
	}
}

