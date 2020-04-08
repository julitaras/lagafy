import moment from "moment";
import CONSTANTS from "../configs/constants";

export const dateParser = date => moment(date).format("L");

export const timeParser = date =>
  moment(date)
    .local()
    .format("LT");

export const isCheckinStarted = time =>
  moment().isBetween(
    moment(time),
    moment(time).add(CONSTANTS.CHECKINTIME, "m"),
    null,
    []
  );

export const isReservationTime = time =>
  moment().isBefore(moment(time))
    ? !moment().isBetween(moment(time).subtract(30, "m"), moment(time))
    : false;

/* eslint-disable */
export const timeParserToUTC = date =>
  moment.utc(new Date(date).toUTCString())._i;
/* eslint-enable */

export const thatTimeHasPassed = time => {
  return moment(time).isAfter(moment());
};
