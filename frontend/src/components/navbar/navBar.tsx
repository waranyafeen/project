import { Link } from "react-router-dom";
import { ModeToggle } from "../modeToggle";
import { HomeIcon } from "lucide-react";

const NavBar = () => {
    return (
        <header className="sticky top-0 z-50 w-full border-b bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
            <div className="flex h-14 items-center w-full px-4 bg-gray-200">
                <div className="flex w-full">
                    <Link className="mr-6  items-center space-x-2 flex" to="/">
                        <HomeIcon />
                        <span className="font-bold ">WARANRAT TOUR</span>
                    </Link>
                    <Link className="mr-6  items-center space-x-2 flex" to="/ticket">
                        <span className="font-bold ">จองตั๋ว</span>
                    </Link>
                </div>
                <div className="flex justify-end w-max items-center gap-2">
                    <ModeToggle />
                </div>
            </div>
        </header>
    );
};

export default NavBar;

