package service

import (
	"errors"
	"fmt"
	"time"

	"hello-run/model"
)

func ObtenerFechasDisponibles(rr model.ReservationClient, servicio string) ([]string, error) {
    hoy := time.Now()
    unAnioDespues := hoy.AddDate(1, 0, 0)
    fechasNoDisponibles := []string{}

    for fecha := hoy; fecha.Before(unAnioDespues); fecha = fecha.AddDate(0, 0, 1) {
        fechaStr := fecha.Format("2006-01-02")

        reservasParaFecha, err := rr.GetAllReservationsByServiceAndDate(servicio, fechaStr)
        if err != nil {
            return nil, err 
        }

        if len(reservasParaFecha) == 4 {
            fechasNoDisponibles = append(fechasNoDisponibles, fechaStr)
			fmt.Printf("Fecha no disponible: %s para el servicio %s", fechaStr, servicio)
        }
    }

    return fechasNoDisponibles, nil
}


func ObtenerHorariosDisponibles(rc model.ReservationClient, servicio string, fecha string) ([]string, error) {

    horariosPredeterminados := []string{
        "9:00 AM", "10:00 AM", "11:00 AM", "2:00 PM", "3:00 PM",
    }

    reservas, err := rc.GetAllReservationsByServiceAndDate(servicio, fecha)
    if err != nil {
        return nil, err
    }

    horariosOcupados := make(map[string]struct{})

    for _, reserva := range reservas {
        horariosOcupados[reserva.Time] = struct{}{}
    }

    horariosDisponibles := []string{}

    for _, horario := range horariosPredeterminados {

        if _, ocupado := horariosOcupados[horario]; !ocupado {
            horariosDisponibles = append(horariosDisponibles, horario)
        }
    }

    return horariosDisponibles, nil
}

func ObtenerMisReservas(rr model.ReservationClient, uc model.UserClient, userEmail string) ([]model.Reservation, error) {
    user, err := uc.UserFirst("email = ?", userEmail)
	if err != nil {
		return nil, errors.New("error trying to find user")
	}


    misReservas, err := rr.GetAllReservationsByUserID(user.ID)
    if err != nil {
        return nil, err // Manejar el error adecuadamente, según tu lógica de negocio.
    }

    fmt.Println("mis reservas:",misReservas);
    

    return misReservas, nil
}

type ReservationParams struct {
	Servicio  string `json:"servicio"`
	Fecha     string `json:"fecha"`
	Horario   string `json:"horario"`
	Ubicacion string `json:"ubicacion"`
	UserEmail string `json:"user_email"`
    TotalPrice float64 `json:"total_price"`
}

func CreateReservation(rr model.ReservationClient, ur model.UserClient, params ReservationParams) (*model.Reservation, error) {
	user, err := ur.UserFirst("email = ?", params.UserEmail)
	if err != nil {
		return nil, errors.New("error trying to find user")
	}

	reserva := &model.Reservation{
		Service:   params.Servicio,
		Date:      params.Fecha,
		Time:      params.Horario,
		Location:  params.Ubicacion,
        TotalPrice: params.TotalPrice,
		UserID:   &user.ID,
	}

	if err := rr.SaveReservation(reserva); err != nil {
		return nil, errors.New("error trying to save reservation")
	}

	return reserva, nil
}

type ReservationCheckParams struct {
	Fecha   string `json:"fecha"`
	Horario string `json:"horario"`
}

func CheckReservation(rr model.ReservationClient, params ReservationCheckParams) (bool, error) {
	return false, errors.New("CheckReservation function not implemented")
}

type ReservationDeleteParams struct { 
    ReservationID uint `json:"reservation_id"`
} 

func DeleteReservation(rr model.ReservationClient, reservationID string) error {
    fmt.Println(reservationID)
    reservation, err := rr.ReservationFirst("id = ?", reservationID)

    if err != nil {
		return errors.New("error trying to find reservation")
	}

    err = rr.DeleteReservation(reservation)

    if err != nil {
		return errors.New("error trying to find reservation")
	}

    return err
}

type EmployeeReservationDoneParams struct {
    ReservationID uint `json:"reservation_id"`
}

func EmployeeReservationDone(rc model.ReservationClient, employeeReservationDoneParams EmployeeReservationDoneParams) (error) {
	reservation, err := rc.ReservationFirst("id = ?", employeeReservationDoneParams.ReservationID)
	if err != nil {
		return errors.New("error trying to find reservation")
	}

	err = rc.UpdateReservation(reservation, &model.Reservation{
		State: "done",
	})
	if err != nil {
		return errors.New("error trying to update reservation")
	}

	return err
}

func GetTotalProfitByMonth(rc model.ReservationClient, month string, year string) (float64, error) {
    profit, err := rc.GetTotalProfitByMonth(month, year)
    if err != nil {
        return 0.0, err
    }

    return profit, nil
}

func GetTotalProfitByYear(rc model.ReservationClient, year string) (float64, error) {
    fmt.Println(year)
    profit, err := rc.GetTotalProfitByYear(year)
    if err != nil {
        return 0.0, err
    }

    return profit, nil
}