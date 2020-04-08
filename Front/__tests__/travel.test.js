/* eslint-disable */

import React from "react";
import { shallow, mount, render } from "enzyme";
import Travel from "../src/components/travels/travelContainer/travel";
import { dateParser, timeParser } from "../src/helpers/dateHelper";

const getTravelDetails = jest.fn();
const toggleTravel = jest.fn();

// defining this.props
const baseProps = {
  id: "1",
  origin: "Dot",
  destination: "Venezuela",
  capacity: "15",
  departure: "2020-02-20 17:30:42",
  arrival: "2020-02-20 18:30:42",
  toggleTravel,
  getTravelDetails
};

describe(" Travel Test", () => {
  let wrapper;
  beforeEach(() => (wrapper = shallow(<Travel {...baseProps} />)));

  it("should call toggleModal and getTravelDetails functions when travel card clicked", () => {
    // Reset info from possible previous calls of these mock functions:
    baseProps.toggleTravel.mockClear();
    // Find the button and call the onClick handler
    wrapper.find(".card").simulate("click");
    // Test to make sure prop functions were called via simulating the button click
    expect(baseProps.toggleTravel).toHaveBeenCalled();
    expect(baseProps.getTravelDetails).toHaveBeenCalled();
  });

  it("Should show capacity, departure and arrival", () => {
    baseProps.toggleTravel.mockClear();
    const li = wrapper.find("li");
    expect(li).toHaveLength(3);
    expect(li.at(0).text()).toContain(`Capacity: ${baseProps.capacity}`);
    expect(li.at(1).text()).toContain(
      `Departure: ${timeParser(baseProps.departure)}`
    );
    expect(li.at(2).text()).toContain(
      `Arrival: ${timeParser(baseProps.arrival)}`
    );
  });

  it("Should show date", () => {
    baseProps.toggleTravel.mockClear();
    const p = wrapper.find("p");
    expect(p).toHaveLength(1);
    expect(p.at(0).text()).toContain(dateParser(baseProps.departure));
  });

  it("Should show origin and destination", () => {
    baseProps.toggleTravel.mockClear();
    const title = wrapper.find(".card-title");
    expect(title).toHaveLength(1);
    expect(title.at(0).text()).toContain(
      `${baseProps.origin} - ${baseProps.destination}`
    );
  });
});
