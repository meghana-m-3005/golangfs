import Greetings from "./greetings";
import Navbar from "./header/navbar";


export default function App(){
  return(
    <>
      <Navbar/>
      <h1><marquee>welcome to REACT</marquee></h1>
      <Greetings/>
    </>
  );
}