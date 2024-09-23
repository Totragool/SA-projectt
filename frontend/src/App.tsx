import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import Payment from './page/Payment.tsx'
import InternetBanking from './page/InternetBanking.tsx'
import Wallet from './page/E-Wallet.tsx'
import Paypal from './page/Paypal.tsx'
import Cards from './page/Cards.tsx'
import './index.css'
import { BrowserRouter as Router, Route, Routes, Link } from 'react-router-dom';

const App: React.FC = () => {
  return (
  <Router>
      <div>
        

        <Routes>
          <Route path="/Payment" element={<Payment />} />
          <Route path="/InternetBanking" element={<InternetBanking />} />
          <Route path="/Wallet" element={<Wallet />} />
          <Route path="/Paypal" element={<Paypal />} />
          <Route path="/Cards" element={<Cards />} />
          <Route path="/" element={<Payment />} /> {/* Make Payment the default route */}
        </Routes>
      </div>
    </Router>
)
}

export default App;

// import React from 'react';
// import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
// import Payment from './page/Payment.tsx';
// import InternetBanking from './page/InternetBanking.tsx';
// import Cards from './page/Cards.tsx';
// import Wallet from './page/E-Wallet.tsx';
// import Paypal from './page/Paypal.tsx';

// const App: React.FC = () => {
//   return (
//     <Router>
//       <Routes>
//         <Route path="/" element={<Payment />} />
//         <Route path="/InternetBanking" element={<InternetBanking />} />
//         <Route path="/Cards" element={<Cards />} />
//         <Route path="/Wallet" element={<Wallet />} />
//         <Route path="/Paypal" element={<Paypal />} />
//       </Routes>
//     </Router>
//   );
// };

// export default App;



