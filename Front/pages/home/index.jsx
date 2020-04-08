import React from "react";
import Home from "../../src/components/home";
import withAuth from "../../src/hoc/withAuth";

//import dynamic from "next/dynamic";
//const Home = dynamic(() => import("../../src/components/home"));

const index = () => {
  return <Home />;
};

export default withAuth(index);
