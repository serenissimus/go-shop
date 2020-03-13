import React from 'react';
import './App.css';
import logo from './logo.svg';
import {Tech} from "./tech/Tech";

export function App() {
    return (
        <div className="app">
            <h2 className="title">go-shop</h2>
            <div className="logo"><img src={logo} height="150px" alt="logo"/></div>
            <div>
                <Tech/>
            </div>
        </div>
    );
}
