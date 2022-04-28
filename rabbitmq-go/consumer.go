package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func ramdom_recursos() string {

	// rand.Seed(time.Now().UnixNano())
	answers := []string{
		"Recurso1",
		"Recurso2",
		"Recurso3",
		"Recurso4",
		"Recurso5",
		"Recurso6",
		"Recurso7",
		"Recurso8",
		"Recurso9",
		"Recurso10",
	}

	return answers[rand.Intn(len(answers))]

}

func ramdom_recursos_utilizar(answers []string) string {

	return answers[rand.Intn(len(answers))]

}

func ramdom_colas_utilizar(answers []string) string {

	return answers[rand.Intn(len(answers))]

}

func main() {

	rand.Seed(time.Now().UnixNano())

	// Numero de recursos que seran asignados (Empleados)
	var num_recursos int
	fmt.Println("Ingrese el numero numero de recursos que seran asignados (Como valor maximo 10):")
	fmt.Scanf("%d", &num_recursos)
	var recursos_utilizar []string
	for i := 0; i < num_recursos; i++ {
		recursos_utilizar = append(recursos_utilizar, ramdom_recursos())

	}
	fmt.Println(recursos_utilizar)

	var colas_utilizar []string
	colas_utilizar = append(colas_utilizar,

		//"Cola1",
		//"Cola2",
		"Cola3",
		//"Cola4",
		//"Cola5",
		"Cola6",
		"Cola7",
		"Cola8",
		//"Cola9",
		//"Cola10",
		//"Cola11",
		//"Cola13",
		"Cola12",
		"Cola14",
		//"Cola15",
	)

	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "No se pudo conectar a RabbitMQ")
	defer connection.Close()

	canal, err := connection.Channel()
	failOnError(err, "Error al abrir un canal")
	defer canal.Close()

	for i := 0; i < len(recursos_utilizar); i++ {
		err = canal.Qos(
			1,     // prefetch count
			0,     // prefetch size
			false, // global
		)
		failOnError(err, "Error al configurar QoS")

		msgs, err := canal.Consume(
			ramdom_colas_utilizar(colas_utilizar), // queue
			"",                                    // consumer
			false,                                 // auto-ack
			false,                                 // exclusive
			false,                                 // no-local
			false,                                 // no-wait
			nil,                                   // args
		)
		failOnError(err, "No se pudo registrar a un consumidor")

		forever := make(chan bool)

		go func() {
			for d := range msgs {
				log.Printf("El paciente : %s, fue atendido por %s y tardo: %s min", d.Body, ramdom_recursos_utilizar(recursos_utilizar), strconv.Itoa(rand.Intn(10-5)+5))
				d.Ack(false)
			}
		}()

		log.Printf(" [*] Esperando mensajes. Para salir presione CTRL+C")
		<-forever
	}

}
