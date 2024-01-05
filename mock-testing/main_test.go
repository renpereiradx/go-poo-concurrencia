package main

import "testing"

// TestGetFullTimeEmployeeByID es una función de prueba para GetFullTimeEmployeeByID.
func TestGetFullTimeEmployeeByID(t *testing.T) {
	// Definimos una tabla de pruebas con diferentes casos de prueba.
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:  1,
			dni: "1",
			// mockFunc es una función que se ejecuta antes de cada prueba para configurar los mocks.
			mockFunc: func() {
				// Aquí estamos reemplazando las funciones GetEmployeeByID y GetPersonByDNI con versiones mock.
				GetEmployeeByID = func(id int) (Employee, error) {
					return Employee{ID: 1, Position: "CEO"}, nil
				}
				GetPersonByDNI = func(dni string) (Person, error) {
					return Person{DNI: "1", Name: "Eduardo", Age: 27}, nil
				}
			},
			// expectedEmployee es el empleado que esperamos obtener de la función GetFullTimeEmployeeByID.
			expectedEmployee: FullTimeEmployee{
				Person:   Person{DNI: "1", Name: "Eduardo", Age: 28},
				Employee: Employee{ID: 1, Position: "CEO"},
			},
		},
	}

	// Guardamos las funciones originales para restaurarlas después de cada prueba.
	originalGetEmployeeByID := GetEmployeeByID
	originalGetPersonByDNI := GetPersonByDNI

	// Iteramos sobre cada caso de prueba en la tabla.
	for _, test := range table {
		// Ejecutamos la función mockFunc para configurar los mocks.
		test.mockFunc()
		// Llamamos a la función que estamos probando.
		ft, err := GetFullTimeEmployeeByID(test.id, test.dni)
		if err != nil {
			// Si hay un error, lo registramos y continuamos con la siguiente prueba.
			t.Errorf("Error getting full time employee by id: %v", err)
		}

		// Comprobamos si el resultado es el esperado.
		if ft.Age != test.expectedEmployee.Age {
			t.Errorf("Expected age %v, got %v", test.expectedEmployee.Person.Age, ft.Person.Age)
		}

		if ft.DNI != test.expectedEmployee.DNI {
			t.Errorf("Expected dni %v, got %v", test.expectedEmployee.Person.DNI, ft.Person.DNI)
		}

		if ft.Name != test.expectedEmployee.Name {
			t.Errorf("Expected name %v, got %v", test.expectedEmployee.Person.Name, ft.Person.Name)
		}

		if ft.Employee.ID != test.expectedEmployee.Employee.ID {
			t.Errorf("Expected employee id %v, got %v", test.expectedEmployee.Employee.ID, ft.Employee.ID)
		}

		if ft.Employee.Position != test.expectedEmployee.Employee.Position {
			t.Errorf("Expected employee position %v, got %v", test.expectedEmployee.Employee.Position, ft.Employee.Position)
		}

		// Restauramos las funciones originales para la siguiente prueba.
		GetEmployeeByID = originalGetEmployeeByID
		GetPersonByDNI = originalGetPersonByDNI
	}
}
