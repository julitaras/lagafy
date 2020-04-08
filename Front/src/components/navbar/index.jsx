import Link from "next/link";
import Logout from "./logout/index";
import React, { useState, Fragment } from "react";
import css from "./index.scss";

const Navbar = ({ isAdmin, currentPage }) => {
  const [showMenu, setShowMenu] = useState(false);

  const displayMenu = () => {
    setShowMenu(!showMenu);
  };

  const linkList = [
    {
      link: "/home",
      title: "Home"
    },
    {
      link: "/travelsManagement",
      title: "Gestión de viajes"
    },
    {
      link: "/reports",
      title: "Reportes"
    }
  ];

  return (
    <Fragment>
      <nav className="navbar navbar-expand-lg navbar-light bg-light">
        <Link href="/">
          <a className="navbar-brand" href="#">
            <i className="icon-Logo-01"></i>
          </a>
        </Link>
        <div className={css.navShowMenuList} onClick={displayMenu}>
          <i className="icon-Menu-00"></i>
        </div>
        {showMenu ? (
          <ul className={css.navMenuList}>
            {isAdmin &&
              linkList
                .filter(l => !l.link.includes(currentPage))
                .map(l => (
                  <Link href={l.link}>
                    <li>{l.title}</li>
                  </Link>
                ))}
            <li>
              <Logout />
            </li>
          </ul>
        ) : (
          ""
        )}
      </nav>
    </Fragment>
  );
};
export default Navbar;
