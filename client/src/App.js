import React from 'react';

import { Routes, Route, BrowserRouter } from 'react-router-dom';
import { PublicRoute, PrivateRoute } from './components/Common';

import { Dashboard, Login, Register } from './pages';

import Auth from './context/AuthContext';

function App() {
    return (
        <BrowserRouter>
            <Auth>
                <Routes>
                    <Route exact path="/login" element={<PublicRoute component={Login} />} />
                    <Route exact path="/register" element={<PublicRoute component={Register} />} />
                    <Route exact path="/" element={<PrivateRoute component={Dashboard} />} />
                </Routes>
            </Auth>
        </BrowserRouter>
    );
}

export default App;
