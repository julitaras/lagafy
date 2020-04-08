import React from "react";
import Link from "next/link";

const index = () => {
  return (
    <div>
      <p>Reportes</p>
      <button>
        <Link href="/">
          <a className="navbar-brand" href="#">
            volver al inicio
          </a>
        </Link>
      </button>
    </div>
  );
};

export default index;
