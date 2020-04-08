import React, { useState } from "react";
import moment from "moment";
import css from "./index.scss";

export const EditTravel = ({
  id,
  haswifi,
  capacity,
  driver,
  departure,
  arrival,
  origin,
  destination,
  status,
  handleDelete,
  handleEdit,
}) => {
  const [isEdit, setIsEdit] = useState(false);
  const [editedTravel, setEditedTravel] = useState({
    id: id,
    haswifi: haswifi,
    capacity: capacity,
    driver: driver,
    departure: moment.utc(departure).format("hh:mm"),
    arrival: moment.utc(arrival).format("hh:mm"),
    origin: origin,
    destination: destination,
    status: status,
    date: moment.utc(departure).format("YYYY-MM-DD"),
  });

  const toggleEdit = () => {
    setIsEdit(!isEdit);
  };
  const handleOnChange = (e) => {
    setEditedTravel({ ...editedTravel, [e.target.name]: e.target.value });
  };
  const handleOnChangeDate = (e) => {
    console.log(editedTravel);
    e.target.value !== null
      ? setEditedTravel({
          ...editedTravel,
          [e.target.name]: e.target.value,
        })
      : null;
  };
  const manageEdit = () =>
    handleEdit(
      moment(
        `${editedTravel.date} ${editedTravel.departure}`,
        "YYYY/MM/DD hh:mm"
      ).toISOString()
    );
  return (
    <div>
      {isEdit ? (
        <div className={css.Card}>
          <div className={`${css.CardForm}`}>
            <div className="form-group form-inline">
              <label htmlFor="origin">Origen:</label>
              <input
                className="form-control"
                id="origin"
                type="text"
                value={editedTravel.origin}
                name="origin"
                onChange={handleOnChange}
                autoFocus
              />
            </div>
            <div className="form-group form-inline">
              <label htmlFor="destination">Destino:</label>
              <input
                className="form-control"
                id="destination"
                type="text"
                value={editedTravel.destination}
                name="destination"
                onChange={handleOnChange}
              />
            </div>
            <div className="form-group form-inline">
              <label htmlFor="date">Fecha:</label>
              <input
                className="form-control"
                id="date"
                type="date"
                value={editedTravel.date}
                name="date"
                onChange={handleOnChangeDate}
                max="2100-12-30"
              />
            </div>
            <div className="form-group form-inline">
              <label htmlFor="departure">Salida:</label>
              <input
                className="form-control"
                id="departure"
                type="time"
                value={editedTravel.departure}
                name="departure"
                onChange={handleOnChangeDate}
                step="600"
              />
            </div>
            <div className="form-group form-inline">
              <label htmlFor="arrival">LLegada:</label>
              <input
                className="form-control"
                id="arrival"
                type="time"
                value={editedTravel.arrival}
                name="arrival"
                onChange={handleOnChangeDate}
              />
            </div>
            <div className="form-group form-inline">
              <label htmlFor="status">Estado:</label>
              <input
                className="form-control"
                id="status"
                type="text"
                value={editedTravel.status}
                name="status"
                onChange={handleOnChange}
              />
            </div>
            <div className="form-group form-inline ">
              <label htmlFor="driver">Chofer:</label>
              <input
                className="form-control"
                id="driver"
                type="text"
                value={editedTravel.driver}
                name="driver"
                onChange={handleOnChange}
              />
            </div>
            <div className="form-group form-inline">
              <label htmlFor="capacity">Capacidad:</label>
              <input
                className="form-control"
                id="capacity"
                type="number"
                min="1"
                max="100"
                value={editedTravel.capacity}
                name="capacity"
                onChange={handleOnChange}
              />
            </div>
            <div>
              <label htmlFor="hasWifi">Wifi:</label>
              <select
                //className="form-control"
                id="hasWifi"
                type="text"
                value={editedTravel.haswifi}
                name="hasWifi"
                onChange={handleOnChange}
              >
                <option>{editedTravel.haswifi ? "Si" : "No"}</option>
                <option>{!editedTravel.haswifi ? "Si" : "No"}</option>
              </select>
            </div>
          </div>
          <p></p>
          <div className={css.ButtonContainer}>
            <button className="btn btn-danger" onClick={toggleEdit}>
              Atras
            </button>
            <button className="btn btn-success" onClick={() => manageEdit()}>
              Aceptar
            </button>
          </div>
        </div>
      ) : (
        <div className={css.Card}>
          <div className={css.CardContent}>
            <p>Origen: {origin}</p>
            <p>Destino: {destination}</p>
            <p>Fecha: {moment.utc(departure).format("DD/MM")}</p>
            <p>Salida: {moment.utc(departure).format("LT")}</p>
            <p>Regreso: {moment.utc(arrival).format("LT")}</p>
            <p>Estado: {status}</p>
            <p>Chofer: {driver}</p>
            <p>Capacidad: {capacity}</p>
            <p>Wifi: {haswifi ? "Si" : "No"}</p>
          </div>
          <div className={css.ButtonContainer}>
            <button className="btn btn-primary" onClick={toggleEdit}>
              Editar
            </button>
            <button className="btn btn-danger" onClick={handleDelete}>
              Borrar
            </button>
          </div>
        </div>
      )}
    </div>
  );
};
export default EditTravel;
