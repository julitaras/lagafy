import React, { useState } from "react";
import TravelContainer from "./travelContainer";
import ReservationContainer from "./reservationContainer";
import Layout from "../layout";
import { userContext } from "../../context";
import withAuth from "../../hoc/withAuth";
import TravelDetails from "./travelContainer/travelDetails";
import css from "../../../style.scss";
import cssTravel from "./index.scss";

const index = () => {
  const { user } = userContext();
  const [selectedTravel, setSelectedTravel] = useState(null);
  const [loadingTravels, setLoadingTravels] = useState(true);
  const [loadingReservations, setLoadingReservations] = useState(true);

  const isLoading = () => loadingTravels || loadingReservations;

  const onClickTravelCard = travel => {
    setSelectedTravel(travel);
    setLoadingTravels(true);
    setLoadingReservations(true);
  };

  return (
    <>
      <Layout currentPage="home">
        {selectedTravel === null ? (
          <div className="container">
            {isLoading() && (
              <div className={css.containerSpinner}>
                <div className="spinner-border text-primary" role="status">
                  <span className="sr-only">Loading...</span>
                </div>
              </div>
            )}

            {!isLoading() && (
              <>
                <h1 className="bold">Hola {user.name}!</h1>
                <p className={cssTravel.textPlane}>A dónde querés ir hoy?</p>
              </>
            )}
            <TravelContainer
              onClickTravelCard={onClickTravelCard}
              loading={loadingTravels}
              setLoading={setLoadingTravels}
              isLoading={isLoading}
            />

            {!isLoading() && <p>Mis Reservas</p>}
            <div className={cssTravel.containerMisReservas}>
              <ReservationContainer
                loading={loadingReservations}
                setLoading={setLoadingReservations}
                isLoading={isLoading}
              />
            </div>
          </div>
        ) : (
          <TravelDetails
            travel={selectedTravel}
            setSelectedTravel={setSelectedTravel}
          />
        )}
      </Layout>
    </>
  );
};

export default withAuth(index);
