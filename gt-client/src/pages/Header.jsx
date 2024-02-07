import { useState } from 'react';
import { Outlet, Link } from 'react-router-dom';


const Header = () => {
  const [current, setCurrent] = useState('h');
  const onClick = (e) => {
    console.log('click ', e);
    console.log(current);
    //setCurrent(e.key);
  };
  return (
    <>
     <div onClick={onClick}>
      <a key="h">
       <Link to="/">Home</Link>
      </a>
      <a key="r">
        <Link to="/register">Register</Link>
      </a>
      <a key="l">
        <Link to="/login">Login</Link>
      </a>
     </div>
     <Outlet />
    </>
   
  )
};
export default Header;