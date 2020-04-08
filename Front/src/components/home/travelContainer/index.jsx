import React, { useState, useEffect } from "react";
import Service from "../../../configs/services";
import TravelCard from "./travelCard";
import { userContext } from "../../../context";

export const TravelContainer = ({
  onClickTravelCard,
  loading,
  setLoading,
  isLoading
}) => {
  const { token } = userContext();
  const [activeTravels, setActiveTravels] = useState([]);

  useEffect(() => {
    getTravels();
  }, []);

  const getTravels = () => {
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
        setLoading(false);
      });
  };

  return (
    <div className="row justify-content-md-center justify-content-left">
      {activeTravels !== null &&
        !isLoading() &&
        activeTravels.map(t => (
          // <EditTravel
          //   key={t.id}
          //   id={t.id}
          //   origin={t.origin}
          //   destination={t.destination}
          //   departure={t.departure}
          //   arrival={t.arrival}
          //   onClickTravelCard={onClickTravelCard}
          //   status={t.status}
          //   driver={t.driver}
          //   capacity={t.capacity}
          //   haswifi={t.haswifi}
          // />
          <TravelCard
            key={t.id}
            id={t.id}
            origin={t.origin}
            destination={t.destination}
            departure={t.departure}
            arrival={t.arrival}
            onClickTravelCard={onClickTravelCard}
          />
        ))}
    </div>
  );
};
export default TravelContainer;
