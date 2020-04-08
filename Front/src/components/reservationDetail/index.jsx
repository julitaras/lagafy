import React, { useEffect, useState } from "react";
import withAuth from "../../hoc/withAuth";
import { useRouter } from "next/router";
import Service from "../../configs/services";
import { userContext } from "../../context";
import {
  timeParser,
  dateParser,
  isCheckinStarted
} from "../../helpers/dateHelper";
import Link from "next/link";
import cssDetail from "./index.scss";
import Layout from "../layout";
import css from "../../../style.scss";

const index = () => {
  const { token, user } = userContext();
  const router = useRouter();
  const { params } = router.query;
  const [travel, setTravel] = useState(null);
  const [errorMsg, setErrorMsg] = useState(null);
  const [loading, setLoading] = useState(false);
  const [isChecked, setIsChecked] = useState(false);
  const [showList, setShowList] = useState(false);

  useEffect(() => {
    getTravel();
  }, []);

  const getTravel = () => {
    setErrorMsg(null);
    setLoading(true);
    Service(token)
      .getTravelDetails(params[0])
      .then(res => {
        if (res.status === 200) {
          return res.json();
        } else {
          setErrorMsg(`Error:`);
        }
      })
      .then(res => {
        if (errorMsg === null) {
          setTravel(res);
        } else {
          setErrorMsg(`${errorMsg} ${res.message}`);
        }
      })
      .then(() => {
        if (errorMsg === null) {
          params[2] === "onboard" ? setIsChecked(true) : setIsChecked(false);
          setLoading(false);
        }
      });
  };

  const handleCheckIn = () => {
    //todo revisar horario
    setLoading(true);
    setErrorMsg(null);
    if (isCheckinStarted(travel.departure)) {
      Service(token)
        .putCheckin(params[1])
        .then(res => {
          if (res.status === 200) {
            return res.json();
          } else {
            setErrorMsg(`Error:`);
          }
        })
        .then(res => {
          if (errorMsg === null) {
            setIsChecked(true);
            // getTravel();
            setTravel({
              ...travel,
              reservations: travel.reservations.map(r =>
                r.passenger.email === user.userName
                  ? { ...r, status: "onboard" }
                  : r
              )
            });
          } else {
            setErrorMsg(`${errorMsg} ${res.message}`);
          }

          setLoading(false);
        });
    } else {
      setErrorMsg("No es horario de checkin");
    }
  };

  const remainingPassengers = travel
    ? travel.reservations.filter(r => r.status === "confirmed")
    : [];

  return (
    <Layout currentPage="home">
      {loading && (
        <div className={css.containerSpinner}>
          <div className="spinner-border text-primary" role="status">
            <span className="sr-only">Loading...</span>
          </div>
        </div>
      )}
      {travel && (
        <div className="container">
          <div className={cssDetail.containerDetail}>
            <Link href="/home">
              <h1 className="bold">
                <i className="icon-Back-00"></i>Check in
              </h1>
            </Link>
            {/* <h1>{`Recorrido ${travel.origin} a ${travel.destination}`}</h1> */}
            <div className={cssDetail.datesDetails}>
              <p>{`Partida: ${timeParser(travel.departure)}`}</p>
              <p>{`Llegada: ${timeParser(travel.arrival)}`}</p>
              <p>{`Fecha: ${dateParser(travel.departure)}`}</p>
              <div className={`${cssDetail.containerCheck}`}>
                <button
                  disabled={isChecked || !isCheckinStarted(travel.departure)}
                  className={`btn ${css.btnSet}`}
                  onClick={handleCheckIn}
                >
                  {isCheckinStarted(travel.departure) ? (
                    <div>
                      <i className="icon-Checkmark-00"></i>
                      Ya me subí
                    </div>
                  ) : (
                    "Fuera de Horario"
                  )}
                </button>

                {errorMsg && (
                  <p className={cssDetail.errorMsg}>
                    {
                      errorMsg // TODO Agregar Estilos
                    }
                  </p>
                )}
              </div>
              <div className={` ${cssDetail.containerList} `}>
                <div className={`${cssDetail.containerChild}`}>
                  <i className="icon-Big-bus-00"></i>
                  {travel.reservations !== undefined && ( //todo sacar esto cuando funcione loading
                    <p>{`Faltan ${remainingPassengers.length} lagashers para irnos!`}</p>
                  )}
                  <button
                    className={`btn ${cssDetail.btnEmpty}`}
                    onClick={() => setShowList(!showList)}
                  >
                    {showList ? "Ocultar lista" : "Ver Lista"}
                  </button>
                  <div className={` text-left ${cssDetail.checkList}`}>
                    <ul>
                      {showList === true &&
                        remainingPassengers.map(t => (
                          <li key={t.passenger.id}>{t.passenger.name}</li>
                        ))}
                    </ul>
                  </div>
                </div>
              </div>
              <div className="text-center">
                <Link href="/home">
                  <button className={`btn ${css.btnEmpty}`}>
                    Volver al incio
                  </button>
                </Link>
              </div>
            </div>
          </div>
        </div>
      )}
    </Layout>
  );
};

export default withAuth(index);
