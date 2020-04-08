import React, { useState } from "react";
import { userDispatch } from "../../../context/";
import Router from "next/router";
import css from "./index.scss";

const Index = () => {
  const [showModal, setShowModal] = React.useState(false);
  const dispatch = userDispatch();

  const logout = () => {
    localStorage.removeItem("lagafy");
    dispatch({
      type: "LOGOUT",
    });
    Router.push("/");
  };

  const submit = () => {
    setShowModal(!showModal);
  };

  return (
    <div>
      <li onClick={submit}>Cerrar sesión</li>
      {showModal ? (
        <div className={css.modal}>
          <form className={css.modalContent}>
            <button className={css.cancelBtn} onClick={submit}>
              X
            </button>
            <div className={css.container}>
              <p className={css.modalDetail}>
                ¿Estás seguro que querés cerrar sesión?
              </p>
              <button className={css.logoutBtn} onClick={logout}>
                Cerrar sesión
              </button>
            </div>
          </form>
        </div>
      ) : (
        ""
      )}
    </div>
  );
};

export default Index;
