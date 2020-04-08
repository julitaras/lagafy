import React from "react";
import dynamic from "next/dynamic";

const Login = dynamic(() => import("../src/components/login"));

const index = () => {
  return <Login />;
};

export default index;
