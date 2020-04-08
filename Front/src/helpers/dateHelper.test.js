import moment from "moment";
import { isCheckinStarted, timeParser, dateParser } from "./dateHelper";

test("isCheckingTime return false if more than 15 mins passed.", () => {
  const res = isCheckinStarted(moment().add(20, "m"));
  expect(res).toBe(false);
});

test("isCheckingTime return true if more than 15 mins passed.", () => {
  const res = isCheckinStarted(moment().subtract(5, "m"));
  expect(res).toBe(true);
});

test("isCheckingTime return false if it's before checkin time.", () => {
  const res = isCheckinStarted(moment().subtract(20, "m"));
  expect(res).toBe(false);
});

test("dateParser return true", () => {
  const res = dateParser("2020-02-28 16:00:00");
  expect(res).toBe("February 28, 2020");
});

// test("dateParser return false", () => {
//   const res = dateParser("2020-04-10 17:30:00");
//   expect(res).toBe("February 28, 2020");
// });

test("timeParser return true", () => {
  const res = timeParser("2020-02-28 16:00:00");
  expect(res).toBe("4:00 PM");
});
