import 'cirrus-ui';
import React from "react";
import {FontAwesomeIcon} from '@fortawesome/react-fontawesome'

import {faHomeAlt} from '@fortawesome/free-solid-svg-icons'


class Header extends React.Component {
    render() {
        return (
            <div className="header header-fixed unselectable header-animated">
                <div className="header-brand">
                    <div className="nav-item no-hover">
                        <h6 className="title uppercase">Expenditure tracker</h6>
                    </div>
                    <div className="nav-item nav-btn" id="header-btn"><span/> <span/> <span/></div>
                </div>
                <div className="header-nav" id="header-menu">
                    <div className="nav-left">
                        <div className="nav-item tex-center">
                            <a href="#">
                                <span className="icon">
                                    <FontAwesomeIcon icon={faHomeAlt}/>
                                </span>
                            </a>
                        </div>
                    </div>
                    <div className="nav-right">
                        <div className="nav-item has-sub toggle-hover" id="dropdown">
                            <a className="nav-dropdown-link"> Menu </a>
                            <ul className="dropdown-menu dropdown-animated" role="menu">
                                <li role="menu-item"><a href="#"> First item </a></li>
                                <li role="menu-item"><a href="#"> Second item </a></li>
                                <li role="menu-item"><a href="#"> Third item </a></li>
                            </ul>

                        </div>
                    </div>
                </div>
            </div>

        )
    }
}

export default Header