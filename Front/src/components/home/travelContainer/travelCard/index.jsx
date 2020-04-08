import React from "react";
import css from "./index.scss";
import {
  dateParser,
  timeParser,
  isReservationTime
} from "../../../../helpers/dateHelper";

export const TravelCard = ({
  origin,
  destination,
  departure,
  arrival,
  id,
  onClickTravelCard
}) => {
  return (
    <div className="col-6 col-lg-2">
      <div
        className={`${css.card} card`}
        onClick={() =>
          onClickTravelCard({
            origin: origin,
            destination: destination,
            departure: departure,
            arrival: arrival,
            id: id
          })
        }
        disabled={!isReservationTime(departure)}
      >
        <div className={`${css.cardContain} card-body text-center`}>
          <p className="card-title semiBold text-uppercase">{`${origin} a ${destination}`}</p>
          {/* <p className="card-text semiBold">{`${dateParser(departure)}`}</p> */}
          <div className="iconContent">
            <i className="icon-Bus-00"></i>
          </div>
          <p
            className={`${css.textDeparture} card-title semiBold`}
          >{`${timeParser(departure)}`}</p>
        </div>
      </div>
    </div>
  );
};
export default TravelCard;
