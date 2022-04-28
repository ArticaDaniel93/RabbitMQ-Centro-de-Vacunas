package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func ramdom_colas() string {

	// rand.Seed(time.Now().UnixNano())
	answers := []string{
		"Cola1",
		"Cola2",
		"Cola3",
		"Cola4",
		"Cola5",
		"Cola6",
		"Cola7",
		"Cola8",
		"Cola9",
		"Cola10",
		"Cola11",
		"Cola12",
		"Cola13",
		"Cola14",
		"Cola15",
	}

	return answers[rand.Intn(len(answers))]

}

func ramdom_colas_utilizar(answers []string) string {

	return answers[rand.Intn(len(answers))]

}

func main() {

	rand.Seed(time.Now().UnixNano())

	//Numero de estaciones (Colas) a utilizar
	var num_estaciones int
	fmt.Println("Ingrese el numero de estaciones a utilizar :")
	fmt.Scanf("%d", &num_estaciones)
	var colas_utilizar []string
	for i := 0; i < num_estaciones; i++ {
		colas_utilizar = append(colas_utilizar, ramdom_colas())

	}
	fmt.Println(colas_utilizar)

	//generar tiempo
	//llamar la funcion de numero aleatorio
	// comparar numero generado > frecuencia //esto en funcion de la tabla **/
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "No se pudo conectar a RabbitMQ")
	defer connection.Close()

	canal, err := connection.Channel()
	failOnError(err, "Error al abrir un canal")
	defer canal.Close()

	var frecuencia_4_7_am float64
	frecuencia_4_7_am = rand.Float64()
	fmt.Println("Frecuencia de 4:30 a 7:30 am: ", frecuencia_4_7_am)

	var frecuencia_7_10_am float64
	frecuencia_7_10_am = rand.Float64()
	fmt.Println("Frecuencia de 7:31 a 10:30 am: ", frecuencia_7_10_am)

	var frecuencia_10_12_am float64
	frecuencia_10_12_am = rand.Float64()
	fmt.Println("Frecuencia de 10:31 a 12:00 am: ", frecuencia_10_12_am)

	var frecuencia_12_1_pm float64
	frecuencia_12_1_pm = rand.Float64()
	fmt.Println("Frecuencia de 12:00 a 1:30 pm: ", frecuencia_12_1_pm)

	var frecuencia_1_6_pm float64
	frecuencia_1_6_pm = rand.Float64()
	fmt.Println("Frecuencia de 1:31 a 6:30 pm: ", frecuencia_1_6_pm)

	var frecuencia_6_8_pm float64
	frecuencia_6_8_pm = rand.Float64()
	fmt.Println("Frecuencia de 6:31 a 8:00 pm: ", frecuencia_6_8_pm)

	// Calculando cuantos pacientes llegan por rangos de horas
	var num_pacientes_4_7_am float64
	num_pacientes_4_7_am = frecuencia_4_7_am * 180.0
	var num_pacientes_4_7_am_casteado int = int(num_pacientes_4_7_am)
	fmt.Println("Numero de pacientes de 4:30 a 7:30 am: ", num_pacientes_4_7_am_casteado)

	var num_pacientes_7_10_am float64
	num_pacientes_7_10_am = frecuencia_7_10_am * 180.0
	var num_pacientes_7_10_am_casteado int = int(num_pacientes_7_10_am)
	fmt.Println("Numero de pacientes de 7:31 a 10:30 am: ", num_pacientes_7_10_am_casteado)

	var num_pacientes_10_12_am float64
	num_pacientes_10_12_am = frecuencia_10_12_am * 180.0
	var num_pacientes_10_12_am_casteado int = int(num_pacientes_10_12_am)
	fmt.Println("Numero de pacientes de 10:31 a 12:00 am: ", num_pacientes_10_12_am_casteado)

	var num_pacientes_12_1_pm float64
	num_pacientes_12_1_pm = frecuencia_12_1_pm * 180.0
	var num_pacientes_12_1_pm_casteado int = int(num_pacientes_12_1_pm)
	fmt.Println("Numero de pacientes de 12:00 a 1:30 pm: ", num_pacientes_12_1_pm_casteado)

	var num_pacientes_1_6_pm float64
	num_pacientes_1_6_pm = frecuencia_1_6_pm * 300.0
	var num_pacientes_1_6_pm_casteado int = int(num_pacientes_1_6_pm)
	fmt.Println("Numero de pacientes de 1:31 a 6:30 pm: ", num_pacientes_1_6_pm_casteado)

	var num_pacientes_6_8_pm float64
	num_pacientes_6_8_pm = frecuencia_6_8_pm * 120.0
	var num_pacientes_6_8_pm_casteado int = int(num_pacientes_6_8_pm)
	fmt.Println("Numero de pacientes de 6:31 a 8:00 pm: ", num_pacientes_6_8_pm_casteado)

	// Generando los numeros aleatorios de los pacientes hacia las colas de manera aleatoria, en horario de 4:30 a 7:30 am
	for i := 0; i < num_pacientes_4_7_am_casteado; i++ {
		var num_pacientes_aleatorio float64
		num_pacientes_aleatorio = rand.Float64()
		if num_pacientes_aleatorio > frecuencia_4_7_am {

			body := "El paciente en horario de 4:30 a 7:30 am se manda a la cola" //bodyFrom(os.Args)
			err = canal.Publish(
				"Estaciones_Vacunas",                  // exchange
				ramdom_colas_utilizar(colas_utilizar), // routing key
				false,                                 // mandatory
				false,
				amqp.Publishing{
					DeliveryMode: amqp.Persistent,
					ContentType:  "text/plain",
					Body:         []byte(body),
				})
			failOnError(err, "Error al publicar un mensaje")
			log.Printf(" [x] Sent %s", body)

		} else {

		}

	}

	// Generando los numeros aleatorios de los pacientes hacia las colas de manera aleatoria, en horario de 7:31 a 10:30 am
	for i := 0; i < num_pacientes_4_7_am_casteado; i++ {
		var num_pacientes_aleatorio float64
		num_pacientes_aleatorio = rand.Float64()
		if num_pacientes_aleatorio > frecuencia_7_10_am {

			body := "El paciente en horario de de 7:31 a 10:30 am se manda a la cola" //bodyFrom(os.Args)
			err = canal.Publish(
				"Estaciones_Vacunas",                  // exchange
				ramdom_colas_utilizar(colas_utilizar), // routing key
				false,                                 // mandatory
				false,
				amqp.Publishing{
					DeliveryMode: amqp.Persistent,
					ContentType:  "text/plain",
					Body:         []byte(body),
				})
			failOnError(err, "Error al publicar un mensaje")
			log.Printf(" [x] Sent %s", body)

		} else {

		}

	}

	// Generando los numeros aleatorios de los pacientes hacia las colas de manera aleatoria, en horario de 10:31 a 12:00 am
	for i := 0; i < num_pacientes_4_7_am_casteado; i++ {
		var num_pacientes_aleatorio float64
		num_pacientes_aleatorio = rand.Float64()
		if num_pacientes_aleatorio > frecuencia_10_12_am {

			body := "El paciente en horario de de 10:31 a 12:00 am se manda a la cola" //bodyFrom(os.Args)
			err = canal.Publish(
				"Estaciones_Vacunas",                  // exchange
				ramdom_colas_utilizar(colas_utilizar), // routing key
				false,                                 // mandatory
				false,
				amqp.Publishing{
					DeliveryMode: amqp.Persistent,
					ContentType:  "text/plain",
					Body:         []byte(body),
				})
			failOnError(err, "Error al publicar un mensaje")
			log.Printf(" [x] Sent %s", body)

		} else {

		}

	}

	// Generando los numeros aleatorios de los pacientes hacia las colas de manera aleatoria, en horario de 12:00 a 1:30 pm
	for i := 0; i < num_pacientes_4_7_am_casteado; i++ {
		var num_pacientes_aleatorio float64
		num_pacientes_aleatorio = rand.Float64()
		if num_pacientes_aleatorio > frecuencia_12_1_pm {

			body := "El paciente en horario de 12:00 a 1:30 pm am se manda a la cola" //bodyFrom(os.Args)
			err = canal.Publish(
				"Estaciones_Vacunas",                  // exchange
				ramdom_colas_utilizar(colas_utilizar), // routing key
				false,                                 // mandatory
				false,
				amqp.Publishing{
					DeliveryMode: amqp.Persistent,
					ContentType:  "text/plain",
					Body:         []byte(body),
				})
			failOnError(err, "Error al publicar un mensaje")
			log.Printf(" [x] Sent %s", body)

		} else {

		}

	}

	// Generando los numeros aleatorios de los pacientes hacia las colas de manera aleatoria, en horario de 1:31 a 6:30 pm
	for i := 0; i < num_pacientes_4_7_am_casteado; i++ {
		var num_pacientes_aleatorio float64
		num_pacientes_aleatorio = rand.Float64()
		if num_pacientes_aleatorio > frecuencia_1_6_pm {

			body := "El paciente en horario de 1:31 a 6:30 pm am se manda a la cola" //bodyFrom(os.Args)
			err = canal.Publish(
				"Estaciones_Vacunas",                  // exchange
				ramdom_colas_utilizar(colas_utilizar), // routing key
				false,                                 // mandatory
				false,
				amqp.Publishing{
					DeliveryMode: amqp.Persistent,
					ContentType:  "text/plain",
					Body:         []byte(body),
				})
			failOnError(err, "Error al publicar un mensaje")
			log.Printf(" [x] Sent %s", body)

		} else {

		}

	}

	// Generando los numeros aleatorios de los pacientes hacia las colas de manera aleatoria, en horario de 6:31 a 8:00 pm
	for i := 0; i < num_pacientes_4_7_am_casteado; i++ {
		var num_pacientes_aleatorio float64
		num_pacientes_aleatorio = rand.Float64()
		if num_pacientes_aleatorio > frecuencia_6_8_pm {

			body := "El paciente en horario de 6:31 a 8:00 pm am se manda a la cola" //bodyFrom(os.Args)
			err = canal.Publish(
				"Estaciones_Vacunas",                  // exchange
				ramdom_colas_utilizar(colas_utilizar), // routing key
				false,                                 // mandatory
				false,
				amqp.Publishing{
					DeliveryMode: amqp.Persistent,
					ContentType:  "text/plain",
					Body:         []byte(body),
				})
			failOnError(err, "Error al publicar un mensaje")
			log.Printf(" [x] Sent %s", body)

		} else {

		}

	}
}
