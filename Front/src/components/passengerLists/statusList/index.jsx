import React, { useState, useContext, useReducer } from "react";
import withAuth from "../../../hoc/withAuth/index";
import userContext from "../../../context/userContext";
import css from "./index.scss";

const statusList = ({status, children, passengers}) => {
  return (
  
        <div className={css.containerLists}>
          <ul className="list-group">
          <li className="list-group-item active">{status}</li>
          {passengers.map((p, i) => (
          <li
            key={i}
            className={
              p.status === "onboard" ? "list-group-item " + css.checked : "list-group-item"
            }
          >
            {p.passenger.name}
          </li>
        ))}
        {children}
      </ul>
    </div>


  );
};

export default statusList;
