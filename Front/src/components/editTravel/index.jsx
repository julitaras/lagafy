import React, { useState } from "react";
import moment from "moment";

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
  handleEdit
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
    date: moment.utc(departure).format("YYYY-MM-DD")
  });

  const toggleEdit = () => {
    setIsEdit(!isEdit);
  };
  const handleOnChange = e => {
    setEditedTravel({ ...editedTravel, [e.target.name]: e.target.value });
  };
  const handleOnChangeDate = e => {
    e.target.value !== null
      ? setEditedTravel({
          ...editedTravel,
          [e.target.name]: e.target.value
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
    <div className="col-6 col-lg-2 card-body text-center">
      {isEdit ? (
        <>
          <input
            type="text"
            value={editedTravel.origin}
            name="origin"
            onChange={handleOnChange}
            autoFocus
          />
          <input
            type="text"
            value={editedTravel.destination}
            name="destination"
            onChange={handleOnChange}
          />
          <input
            type="date"
            value={editedTravel.date}
            name="date"
            onChange={handleOnChangeDate}
            max="2100-12-30"
          />
          <input
            type="time"
            value={editedTravel.departure}
            name="departure"
            onChange={handleOnChangeDate}
            step="600"
          />
          <input
            type="time"
            value={editedTravel.arrival}
            name="arrival"
            onChange={handleOnChangeDate}
          />
          <input
            type="text"
            value={editedTravel.status}
            name="status"
            onChange={handleOnChange}
          />
          <input
            type="text"
            value={editedTravel.driver}
            name="driver"
            onChange={handleOnChange}
          />
          <input
            type="number"
            min="1"
            value={editedTravel.capacity}
            name="capacity"
            onChange={handleOnChange}
          />
          <input
            type="text"
            value={editedTravel.haswifi}
            name="capacity"
            onChange={handleOnChange}
          />
          <button onClick={() => manageEdit()}>Confirmar</button>
          <button onClick={toggleEdit}>Cancelar</button>
        </>
      ) : (
        <>
          <p>{origin}</p>
          <p>{destination}</p>
          <p>{moment.utc(departure).format("DD/MM")}</p>
          <p>{moment.utc(departure).format("LT")}</p>
          <p>{moment.utc(arrival).format("LT")}</p>
          <p>{status}</p>
          <p>{driver}</p>
          <p>{capacity}</p>
          <p>{haswifi}</p>
          <button onClick={toggleEdit}>Editar</button>
          <button onClick={handleDelete}>Borrar</button>
        </>
      )}
    </div>
  );
};
export default EditTravel;
