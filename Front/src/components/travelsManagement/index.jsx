import React, { useState, useEffect } from "react";
import { userContext } from "../../context";
import EditTravel from "./editTravel";
import RegisterTravel from "./registerTravels";
import Layout from "../layout";
import Service from "../../configs/services";
import style from "../../../style.scss";
import css from "./index.scss";

const index = () => {
  const { token } = userContext();
  const [travels, setTravels] = useState([]);
  const [loading, setLoading] = useState(true);
  const [showTravels, setShowTravels] = useState(true);

  useEffect(() => {
    getTravels();
  }, []);

  const getTravels = () => {
    Service(token)
      .getCurrentTravels()
      .then((res) => {
        if (res.status === 200) {
          return res.json();
        } else {
          console.log("failed to call travels/current");
        }
      })
      .then((res) => {
        // TODO Validar que la respuesta no devuelva token invalido
        setTravels(res);
        setLoading(false);
      })
      .catch((err) => console.log(err));
  };

  return (
    <Layout currentPage="travelsManagement">
      <div className="container">
        <h1 className="bold">Gestión de viajes</h1>
        <div>
          {loading && (
            <div className={style.containerSpinner}>
              <div className="spinner-border text-primary" role="status">
                <span className="sr-only">Loading...</span>
              </div>
            </div>
          )}
          {!loading && (
            <div className="row justify-content-md-center">
              <div className={css.Header}>
                <button
                  className={`${
                    showTravels ? "btn btn-primary" : "btn btn-danger"
                  } ${css.NewTravel}`}
                  onClick={() => {
                    setShowTravels(!showTravels);
                  }}
                >
                  {showTravels ? "Nuevo viaje" : "Cancelar"}
                </button>
              </div>
            </div>
          )}
          <div className="row justify-content-md-center">
            <div className={css.Container}>
              {!showTravels && <RegisterTravel />}
              {!loading &&
                showTravels &&
                travels.map((t) => {
                  return (
                    <EditTravel
                      key={t.id}
                      id={t.id}
                      origin={t.origin}
                      destination={t.destination}
                      departure={t.departure}
                      arrival={t.arrival}
                      status={t.status}
                      driver={t.driver}
                      capacity={t.capacity}
                      haswifi={t.haswifi}
                    />
                  );
                })}
            </div>
          </div>
        </div>
      </div>
    </Layout>
  );
};

export default index;
