import React from "react";
import TravelsManagement from "../../src/components/travelsManagement";
import withAuth from "../../src/hoc/withAuth";

const index = () => {
  return <TravelsManagement />;
};

export default withAuth(index);
