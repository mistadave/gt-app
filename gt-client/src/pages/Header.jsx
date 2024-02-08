import { useState } from "react";
import {
  NavigationMenu,
  NavigationMenuContent,
  NavigationMenuItem,
  NavigationMenuLink,
  NavigationMenuList,
  NavigationMenuTrigger,
  navigationMenuTriggerStyle,
} from "../components/ui/navigation-menu";
import { cn } from "@/lib/utils";
import { Link } from "react-router-dom";

const Header = () => {
  const [current, setCurrent] = useState("h");
  const onClick = (e) => {
    console.log("click ", e);
    console.log(current);
    //setCurrent(e.key);
  };

  return (
    <>
      <NavigationMenu>
        <NavigationMenuList>
          <Link to="/">
            <NavigationMenuLink className={navigationMenuTriggerStyle()}>
              Home
            </NavigationMenuLink>
          </Link>
          <Link to="/games">
            <NavigationMenuLink className={navigationMenuTriggerStyle()}>
              Games
            </NavigationMenuLink>
          </Link>
          <NavigationMenuItem>
            <NavigationMenuTrigger>Auth</NavigationMenuTrigger>
            <NavigationMenuContent>
              <ul className="grid w-[400px] gap-3 p-4 md:w-[500px] md:grid-cols-2 lg:w-[600px] ">
                <li>
                  <Link to="/login">
                    <NavigationMenuLink
                      className={cn(
                        "block select-none space-y-1 rounded-md p-3 leading-none no-underline outline-none transition-colors hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground"
                      )}
                    >
                      <div className="text-sm font-medium leading-none">
                        Login
                      </div>
                      <p className="line-clamp-2 text-sm leading-snug text-muted-foreground">
                        Login to your account
                      </p>
                    </NavigationMenuLink>
                  </Link>
                </li>
                <li>
                  <Link to="/register">
                    <NavigationMenuLink
                      className={cn(
                        "block select-none space-y-1 rounded-md p-3 leading-none no-underline outline-none transition-colors hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground"
                      )}
                    >
                      <div className="text-sm font-medium leading-none">
                        Register
                      </div>
                      <p className="line-clamp-2 text-sm leading-snug text-muted-foreground">
                        Register new account
                      </p>
                    </NavigationMenuLink>
                  </Link>
                </li>
              </ul>
            </NavigationMenuContent>
          </NavigationMenuItem>
        </NavigationMenuList>
      </NavigationMenu>
    </>
  );
};

export default Header;
