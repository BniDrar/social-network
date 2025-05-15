'use client'
import React, { useState } from 'react'
import { BiMenuAltLeft, BiMenuAltRight } from "react-icons/bi"

export default function ShowOnMobile() {
      if (typeof window !== 'undefined' && window.resizeTo < 768) {
            return (
                  <div className="show-on-mobile">
                        <span className="show-lift-sidebar" ><BiMenuAltLeft /></span>
                        <span className="show-right-sidebar"><BiMenuAltRight /></span>
                  </div>
            );
      }
      const [leftSidebar, setLeftSidebar] = useState(false);
      const [rightSidebar, setRightSidebar] = useState(false);
      const [leftSdebar, setLeftSdebar] = useState(null);
      const [rightSdebar, setRightSdebar] = useState(null);
      const [main, setMain] = useState(null);

      React.useEffect(() => {
            setLeftSdebar(document.querySelector('#left-sidebar'));
            setRightSdebar(document.querySelector('#right-sidebar'));
            setMain(document.querySelector('main'));
      }, []);

      const showLeftSidebar = () => {
            if (leftSidebar && leftSdebar) {
                  setLeftSidebar(false);
                  leftSdebar?.classList.remove('show');
            } else {
                  setLeftSidebar(true);
                  leftSdebar?.classList.add('show');
            }
      }

      const showRightSidebar = () => {
            if (rightSidebar && rightSdebar) {
                  setRightSidebar(false);
                  rightSdebar?.classList.remove('show');
            } else {
                  setRightSidebar(true);
                  rightSdebar?.classList.add('show');
            }
      }
      main?.addEventListener('click', () => {
            setLeftSidebar(false);
            setRightSidebar(false);
            leftSdebar?.classList.remove('show');
            rightSdebar?.classList.remove('show');
      })



      return (
            <div className="show-on-mobile">
                  <span className="show-lift-sidebar" onClick={showLeftSidebar}><BiMenuAltLeft /></span>
                  <span className="show-right-sidebar" onClick={showRightSidebar}><BiMenuAltRight /></span>
            </div>
      )
}