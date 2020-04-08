import moment from "moment";
import { timeParserToUTC } from "./dateHelper";

export const isUserLogged = user => !(user === {} || user === null);
export const getTokenFromLocalStorage = () =>
  JSON.parse(localStorage.getItem("lagafy"));
export const setTokenInLocalStorage = accessToken =>
  localStorage.setItem("lagafy", JSON.stringify(accessToken)) || null;
export const isTokenExpired = () => {
  const accessToken = getTokenFromLocalStorage();
  const tokenExpiration = timeParserToUTC(accessToken.expiresOn);

  return moment(tokenExpiration).isBefore(moment());
};
