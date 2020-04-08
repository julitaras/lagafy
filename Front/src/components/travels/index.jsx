import React, { useState, useContext, useEffect, Fragment } from "react";
import userContext from "../../context/userContext";
import Layout from "../layout";
import Travels from "./travelContainer";
import PassengerLists from "../passengerLists";
import withAuth from "../../hoc/withAuth";
import css from "../../../style.scss";
import Service from "../../configs/services.js";

const Index = () => {
  const { user, token } = useContext(userContext);
  const [showTravel, setShowTravel] = useState(false);
  const [loading, setLoading] = useState(true);
  const [selectedTravel, setSelectedTravel] = useState(null);
  const [activeTravels, setActiveTravels] = useState(null);
  const [travelDetails, setTravelDetails] = useState(null);

  const handleShowTravel = id => {
    setShowTravel(!showTravel);
    setSelectedTravel(id);
  };

  const getTravelDetails = id => {
    setLoading(true);
    if (id !== null && id !== undefined) {
      Service(token)
        .getTravelDetails(id)
        .then(res => {
          if (res.status === 200) {
            return res.json();
          } else {
            console.log("failed to call travels/id");
          }
        })
        .then(res => {
          setTravelDetails(res);
          setLoading(false);
        });
    }
  };

  useEffect(() => {
    getTravels();
  }, [showTravel]);

  const getTravels = () => {
    setLoading(true);
    Service(token)
      .getCurrentTravels()
      .then(res => {
        if (res.status === 200) {
          return res.json();
        } else {
          console.log("failed to call travels/current");
        }
      })
      .then(res => {
        // TODO Validar que la respuesta no devuelva token invalido
        setActiveTravels(res);
        setTimeout(() => {
          setLoading(false);
        }, 800);
      });
  };

  return (
    <Fragment>
      {loading ? (
        <div className={css.containerSpinner}>
          <div className="spinner-border text-primary" role="status">
            <span className="sr-only">Loading...</span>
          </div>
        </div>
      ) : (
        <Layout currentPage="home">
          <h1 className="text-center">Hello {user.displayName}</h1>

          {showTravel && travelDetails ? (
            <PassengerLists
              toggleTravel={handleShowTravel}
              selectedTravel={selectedTravel}
              list={travelDetails}
              setLoading={setLoading}
            />
          ) : (
            <Fragment>
              {!activeTravels.length > 0 ? (
                <div className="container">
                  <div className={css.NoTravels}>
                    <h3 className="text-center">No trips available</h3>
                  </div>
                </div>
              ) : (
                <Travels
                  toggleTravel={handleShowTravel}
                  activeTravels={activeTravels}
                  getTravelDetails={getTravelDetails}
                />
              )}
            </Fragment>
          )}
        </Layout>
      )}
    </Fragment>
  );
};

export default withAuth(Index);
