// import { StrictMode } from 'react'
// import { createRoot } from 'react-dom/client'
// import Payment from './Payment.tsx'
// import InternetBanking from './InternetBanking.tsx'
// import Wallet from './E-Wallet.tsx'
// import Paypal from './Paypal.tsx'
// import Cards from './Cards.tsx'
// import './index.css'
// import { BrowserRouter as Router, Route, Routes, Link } from 'react-router-dom';

// const App: React.FC = () => {
//   return (
//   <Router>
//       <div>
        

//         <Routes>
//           <Route path="/Payment" element={<Payment />} />
//           <Route path="/InternetBanking" element={<InternetBanking />} />
//           <Route path="/Wallet" element={<Wallet />} />
//           <Route path="/Paypal" element={<Paypal />} />
//           <Route path="/Cards" element={<Cards />} />
//           <Route path="/" element={<Payment />} /> {/* Make Payment the default route */}
//         </Routes>
//       </div>
//     </Router>
// )
// }

// export default App;

import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Payment from './Payment';
import InternetBanking from './InternetBanking';
import Cards from './Cards';
import Wallet from './E-Wallet';
import Paypal from './Paypal';

const App: React.FC = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Payment />} />
        <Route path="/InternetBanking" element={<InternetBanking />} />
        <Route path="/Cards" element={<Cards />} />
        <Route path="/Wallet" element={<Wallet />} />
        <Route path="/Paypal" element={<Paypal />} />
      </Routes>
    </Router>
  );
};

export default App;