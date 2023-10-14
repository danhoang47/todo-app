
import { faRightFromBracket } from "@fortawesome/free-solid-svg-icons";
import { faCircleQuestion } from "@fortawesome/free-regular-svg-icons";

import { Category } from "../category";

import "./NavbarFooter.styles.scss"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

const footerOptions = [
    {
        id: 1,
        icon: faCircleQuestion,
        title: 'Help'
    },
    {
        id: 2,
        icon: faRightFromBracket,
        title: 'Logout'
    }
]

function NavbarFooter() {
    return (  
        <Category 
            className="footer" 
            data={footerOptions}
            render={(option) => (
                <div key={option.id} className="categoryItem" tabIndex={0}>
					<span><FontAwesomeIcon className="icon" icon={option.icon} /></span>
					<p className="categoryItemTitle">{option.title}</p>
				</div>
            )}
        />
    );
}

export default NavbarFooter;