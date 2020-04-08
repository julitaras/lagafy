import React from "react";
import Link from "next/link";
import css from "./index.scss";

const RegisterTravel = ({ handleAddTravel }) => {
  const validateNewTravel = () => {
    // console.log('fecha',)
  };

  return (
    <form className={css.formRegister}>
      <div className="form-group">
        <label htmlFor="formGroupInput">Fecha</label>
        <input type="date" name="fecha" />
      </div>
      <div className="form-group">
        <label htmlFor="formGroupInput">Origen</label>
        <select>
          <option value="">Lagash - Caballito</option>
          <option value="">Lagash - Parque Patricios</option>
          <option value="">Meli - Polo Dot</option>
          <option value="">Meli - Tesla</option>
        </select>
      </div>
      <div className="form-group">
        <label htmlFor="destino">Destino</label>
        <select>
          <option value="">Lagash - Caballito</option>
          <option value="">Lagash - Parque Patricios</option>
          <option value="">Meli - Polo Dot</option>
          <option value="">Meli - Tesla</option>
        </select>
      </div>
      <div className="form-group">
        <label htmlFor="formGroupInput">Horario de partida</label>
        <input
          type="time"
          className="form-control"
          placeholder="Ingrese el horario de partida"
        />
      </div>
      <div className="form-group">
        <label htmlFor="formGroupInput">Horario de llegada</label>
        <input
          type="time"
          className="form-control"
          placeholder="Ingrese el horario de llegada"
        />
      </div>
      <div className="form-group">
        <label htmlFor="formGroupInput">Capacidad</label>
        <input
          type="number"
          className="form-control"
          placeholder="Ingrese la cantidad permitida de pasajeros abordo"
        />
      </div>
      <div className="form-group">
        <label htmlFor="formGroupInput">Conductor</label>
        <input
          type="text"
          className="form-control"
          placeholder="Ingrese el nombre del conductor"
        />
      </div>
      <div className="form-group">
        <label htmlFor="wifi">Tiene wifi:</label>
        <select id="wifi">
          <option value="si">Si</option>
          <option value="no">No</option>
        </select>
      </div>
      <div className="form-group">
        <button onClick={validateNewTravel} type="submit">
          Agregar
        </button>
      </div>
    </form>
  );
};

export default RegisterTravel;
