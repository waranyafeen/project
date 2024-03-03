import { Link } from "react-router-dom";
import { HomeIcon } from "lucide-react";
import { ModeToggle } from "@/components/ui/mode-toggle";
//import Home from "@/page/home";
//import { LogIn } from 'lucide-react';

const NavBar = () => {
    return (
        <div>
            <div className="flex h-12 px-4 items-center">
                <div className="flex w-full">
                    <Link className="mr-6 items-center space-x-2 flex" to="/">
                        <HomeIcon/><span className="font-bold ">WARANRAT TOUR</span>
                    </Link>
                    <Link className="mr-6  items-center space-x-2 flex" to="/ticket">
                        <span className="font-bold ">จองตั๋ว</span>
                    </Link>
                </div>
                <div className="flex justify-end w-max items-center gap-4">
                    <Link to="/login" className="font-bold">
                        {/* <LogIn/> */}
                        LOGIN
                    </Link>
                    <ModeToggle />
                </div>
            </div>
        </div>
    );
};

export default NavBar;
