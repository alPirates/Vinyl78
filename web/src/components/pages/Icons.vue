<template lang="pug">
  v-layout(row, wrap)
    v-flex(xs12)
      v-layout(row, wrap).hide-mobile
        v-flex(xs12 sm4)
          router-link(tag="div" :to="'/category/' + thumbLinks.brand")
            .thumb(@mouseover="hover.brand = true", @mouseleave="hover.brand = false", :class="{hovering: hover.brand}")
              svg(type="image/svg+xml" class="icon-brand svg-icon" width="128" height="128" viewBox="0 0 512 512")
                <path class="fil0" d="M423.14 359.03c-8.77,0 -15.92,7.14 -15.92,15.92 0,8.78 7.15,15.92 15.92,15.92 8.78,0 15.92,-7.14 15.92,-15.92 0,-8.78 -7.14,-15.92 -15.92,-15.92z"/>
                <path class="fil0" d="M106.94 359.03c-8.77,0 -15.92,7.14 -15.92,15.92 0,8.78 7.15,15.92 15.92,15.92 8.78,0 15.92,-7.14 15.92,-15.92 0,-8.78 -7.14,-15.92 -15.92,-15.92z"/>
                <path class="fil0" d="M489.89 255.07l-60.57 -26.92 -44.33 -96.06c-5.64,-12.21 -17.96,-20.09 -31.41,-20.09l-325.69 0c-14.22,0 -25.8,11.57 -25.8,25.8l0 219.58c0,14.23 11.58,25.81 25.8,25.81l27.56 0c3.96,24.85 25.54,43.91 51.49,43.91 25.96,0 47.54,-19.06 51.5,-43.91l213.21 0c3.96,24.85 25.54,43.91 51.49,43.91 25.96,0 47.54,-19.06 51.5,-43.91l9.99 0c14.22,0 25.8,-11.58 25.8,-25.81l0 -70.7c0,-13.66 -8.06,-26.06 -20.54,-31.61zm-104.44 -82.68l24.82 53.79 -75.51 0 0 -53.8 50.69 0.01zm-278.51 238.24c-19.67,0 -35.68,-16 -35.68,-35.68 0,-19.67 16.01,-35.68 35.68,-35.68 19.68,0 35.69,16.01 35.69,35.68 0,19.68 -16.01,35.68 -35.69,35.68zm316.2 0c-19.67,0 -35.68,-16 -35.68,-35.68 0,-19.67 16.01,-35.68 35.68,-35.68 19.68,0 35.69,16.01 35.69,35.68 0,19.68 -16.01,35.68 -35.69,35.68zm70.82 -79.05l-9.33 0c-4.55,0 -8.24,3.69 -8.24,8.24 0,4.55 3.69,8.23 8.24,8.23l9.33 0 0 9.33c0,5.15 -4.19,9.34 -9.33,9.34l-9.99 0c-3.96,-24.86 -25.54,-43.92 -51.5,-43.92 -25.95,0 -47.53,19.06 -51.49,43.92l-213.21 0c-3.96,-24.86 -25.54,-43.92 -51.5,-43.92 -25.95,0 -47.53,19.06 -51.49,43.92l-27.56 0c-5.14,0 -9.33,-4.19 -9.33,-9.34l0 -9.33 26.9 0c4.55,0 8.23,-3.68 8.23,-8.23 0,-4.55 -3.68,-8.24 -8.23,-8.24l-26.9 0 0 -193.78c0,-5.15 4.19,-9.33 9.33,-9.33l325.69 0c7.05,0 13.5,4.13 16.45,10.52l7.81 16.93 -77.66 0c-9.39,0 -17.02,7.63 -17.02,17.01l0 35.14c0,19.07 15.51,34.58 34.58,34.58l103.66 0 51.13 22.73 -1.29 0.79c-7.66,4.64 -12.41,13.09 -12.41,22.05l0 16.47c0,4.54 3.68,8.23 8.23,8.23l26.9 0 0 18.66zm-175.67 -159.19l0 53.79 -0.55 0c-9.99,0 -18.11,-8.12 -18.11,-18.11l0 -35.14c0,-0.3 0.24,-0.55 0.55,-0.55l18.11 0.01zm175.67 124.06l-18.66 0 0 -8.23c0,-3.24 1.72,-6.29 4.48,-7.98l9.46 -5.74c2.98,3.27 4.72,7.59 4.72,12.17l0 9.78z"/>
                <path class="fil0" d="M361.66 331.58l-193.23 0c-4.55,0 -8.24,3.69 -8.24,8.24 0,4.55 3.69,8.23 8.24,8.23l193.23 0c4.55,0 8.24,-3.68 8.24,-8.23 0,-4.55 -3.69,-8.24 -8.24,-8.24z"/>
                <path class="fil1" d="M123.26 182.69l-9.15 -21.92 -11.36 4.75 9.19 22.03c1.01,2.41 2.25,4.03 3.73,4.86 1.46,0.82 3.15,0.84 5.04,0.05 1.87,-0.78 3.03,-1.98 3.47,-3.6 0.43,-1.63 0.12,-3.68 -0.92,-6.17zm-17.38 10.24l-10.2 -24.46 -15.89 6.62 10.21 24.47c1.1,2.65 2.65,4.41 4.68,5.29 1.99,0.86 4.27,0.78 6.8,-0.28 2.5,-1.04 4.12,-2.58 4.91,-4.61 0.76,-2.04 0.6,-4.38 -0.51,-7.03zm-36.19 -22.47l48.24 -20.13 12.38 29.69c2.05,4.91 2.6,9.06 1.63,12.43 -0.96,3.38 -3.41,5.87 -7.31,7.5 -1.97,0.82 -3.87,1.16 -5.73,1.01 -1.86,-0.16 -3.58,-0.8 -5.16,-1.91 -0.38,3.31 -1.37,6.11 -2.99,8.37 -1.59,2.28 -3.77,3.98 -6.47,5.1 -4.5,1.88 -8.67,1.85 -12.51,-0.06 -3.85,-1.94 -6.87,-5.55 -9.08,-10.85l-13 -31.15zm61.81 51.35l20.41 18.95c2.07,1.92 4.06,2.95 5.99,3.05 1.93,0.11 3.68,-0.67 5.25,-2.35 1.55,-1.68 2.2,-3.48 1.92,-5.37 -0.26,-1.9 -1.43,-3.81 -3.5,-5.73l-20.41 -18.95 -9.66 10.4zm-5.22 5.62l-15.56 16.75 -5.98 -5.56 35.56 -38.29 26.4 24.5c3.92,3.65 6.11,7.45 6.58,11.41 0.48,3.95 -0.86,7.6 -3.98,10.96 -2.52,2.71 -5.59,4.27 -9.2,4.62 -3.61,0.37 -7.03,-0.52 -10.27,-2.7l-4.95 26.78 -6.52 -6.06 4.66 -26.87 -16.74 -15.54zm52 43.63l18.86 7.99 -1.54 -22.46 -17.32 14.47zm-25.37 11.22l44.97 -37.85 4.75 2.01 4.31 58.72 -8.08 -3.42 -1.08 -14.92 -25.47 -10.78 -11.42 9.62 -7.98 -3.38zm71.44 26.91l7.74 -51.82 4.25 0.64 27.46 41.33 5.45 -36.41 7.95 1.19 -7.75 51.82 -4.24 -0.64 -27.49 -41.2 -5.42 36.28 -7.95 -1.19zm89.63 -45.13c8.53,-0.77 15.61,1.04 21.29,5.43 5.65,4.4 8.83,10.48 9.54,18.26 0.72,7.86 -1.28,14.37 -5.96,19.55 -4.68,5.16 -11.39,8.15 -20.13,8.94l-23.1 2.1 -4.75 -52.18 23.11 -2.1zm0.8 7.5l-15.09 1.37 3.38 37.16 15.1 -1.37c6.19,-0.56 10.92,-2.67 14.16,-6.31 3.24,-3.64 4.61,-8.33 4.09,-14.02 -0.5,-5.51 -2.75,-9.82 -6.73,-12.95 -3.99,-3.13 -8.94,-4.42 -14.91,-3.88z"/>
              h2.thumb-text.text-xs-center БРЕНДИРОВАНИЕ АВТОМОБИЛЕЙ
        v-flex(xs12 sm4)
          router-link(tag="div" :to="'/category/' + thumbLinks.vinyl")
            .thumb(@mouseover="hover.vinyl = true", @mouseleave="hover.vinyl = false", :class="{'vinyl-hover': hover.vinyl, hovering: hover.vinyl}")
              <svg  class="svg-icon icon-vinyl" xmlns="http://www.w3.org/2000/svg" xml:space="preserve" width="128" height="128" version="1.1" style="shape-rendering:geometricPrecision; text-rendering:geometricPrecision; image-rendering:optimizeQuality; fill-rule:evenodd; clip-rule:evenodd" viewBox="0 0 512 512" xmlns:xlink="http://www.w3.org/1999/xlink">
              <metadata id="CorelCorpID_0Corel-Layer"/>
              <path class="fil0 str0" d="M68.82 234.26l230.05 -92.94 73.92 182.96 -230.05 92.95c-50.52,20.41 -108.03,-4 -128.44,-54.52 -20.41,-50.53 4,-108.03 54.52,-128.45zm178.86 140.57c-50.53,20.41 -108.03,-4 -128.44,-54.52 -20.42,-50.53 3.99,-108.03 54.51,-128.44"/>
              <polygon class="fil1 str1" points="355.12,291.44 343.96,263.83 354.68,221.21 317.39,198.06 306.8,171.84 307.83,171.43 358.72,204.54 372.5,145.3 392.7,137.14 373.21,213.72 447.88,260.01 427.6,268.2 369.32,230.27 "/>
              <polygon class="fil1 str1" points="346.76,300.86 333.29,306.3 339.3,282.39 "/>
              <polygon class="fil1 str1" points="301.15,187.98 287.63,179.59 296.34,176.07 "/>
              <path class="fil1 str1" d="M151.7 251.74l-5.98 -14.81 16.19 -6.55 5.99 14.82 -16.2 6.54zm-16.44 -10.59l5.98 14.82 -14.63 5.91 5.41 13.39 -16.59 6.71 -5.41 -13.4 -12.71 5.14 -5.99 -14.82 43.94 -17.75zm-48.41 36.79l-17.77 7.18 -5.99 -14.81 17.77 -7.18 5.99 14.81zm67.67 53.03l-16.59 6.71 -18.27 -45.24 16.59 -6.71 18.27 45.24zm69.92 -64.27l5.99 14.82 -6.83 2.76 6.53 16.15 -16.54 6.68 -18.27 -45.24 16.53 -6.68 5.76 14.27 6.83 -2.76zm29.03 71.85l-23.36 9.44 -12.29 -30.42 16.53 -6.68 6.31 15.6 6.82 -2.76 5.99 14.82zm-111.31 9.59l16.59 -6.71 12.29 30.43 -16.59 6.7 -12.29 -30.42zm60.89 -134.38l5.99 14.82 -6.83 2.76 5.41 13.39 -16.53 6.68 -11.4 -28.21 23.36 -9.44zm56.36 -22.77l10.44 -4.22 5.99 14.82 -10.44 4.22 -5.99 -14.82zm21.39 52.94l5.67 -2.29 5.99 14.81 -5.67 2.29 -5.99 -14.81zm23.05 57.03l13.34 -5.39 5.99 14.82 -13.35 5.39 -5.98 -14.82zm-10.46 4.23l5.98 14.82 -12.43 5.02 -5.98 -14.82 12.43 -5.02zm-23.05 -57.04l5.99 14.82 -12.43 5.02 -5.99 -14.81 12.43 -5.03zm-27.83 -33.1l-5.99 -14.81 12.43 -5.03 5.99 14.82 -12.43 5.02zm-23.01 9.3l-5.99 -14.82 12.55 -5.07 5.99 14.82 -12.55 5.07zm21.39 52.94l-5.99 -14.82 12.55 -5.07 5.99 14.82 -12.55 5.07zm23.04 57.04l-5.99 -14.82 12.56 -5.07 5.98 14.82 -12.55 5.07z"/>
              <polygon class="fil1 str1" points="506.95,108.22 465.66,124.9 510.09,234.87 493.5,241.58 449.07,131.6 408.13,148.14 402.14,133.32 500.96,93.4 "/>
              </svg>
              h2.thumb-text.text-xs-center ВИНИЛОВЫЕ НАКЛЕЙКИ
        v-flex(xs12 sm4)
          router-link(tag="div" :to="'/category/' + thumbLinks.decor")
            .thumb(@mouseover="hover.decor = true", @mouseleave="hover.decor = false", :class="{hovering: hover.decor}")
              <svg class="icon-brand svg-icon" xmlns="http://www.w3.org/2000/svg" xml:space="preserve" width="128" height="128" version="1.1" style="shape-rendering:geometricPrecision; text-rendering:geometricPrecision; image-rendering:optimizeQuality; fill-rule:evenodd; clip-rule:evenodd" viewBox="0 0 2167.46 2167.46" xmlns:xlink="http://www.w3.org/1999/xlink">
                <g id="Слой_x0020_1">
                <metadata id="CorelCorpID_0Corel-Layer"/>
                <path class="fil0 str0" d="M2036.43 638.53l41.32 66.29c-104.04,-1.18 -1880.24,-2.9 -1991.63,0l41.32 -68.02 193.87 -131.7c6.31,-4.58 19,-7.77 31.7,-7.77l1459.75 0.37c12.74,0 24.55,2.82 30.9,7.4l192.77 133.43 0 0zm62.01 -1.64l47.79 83.61c2.69,4.75 4.12,10.09 4.12,15.55 0,102.45 -55.36,173.74 -133.72,213.34l-0.04 959.14 35.61 0c17.27,0 31.27,14 31.27,31.27 0,17.28 -14,31.28 -31.27,31.28l-1939.43 0c-17.28,0 -31.28,-14 -31.28,-31.28 0,-17.27 14,-31.27 31.28,-31.27l35.6 0 1.43 -957.42c-80.08,-32.75 -135.19,-112.61 -135.19,-215.06 0,-5.46 1.43,-10.8 4.12,-15.55l47.76 -83.61 200.51 -144.61c17.4,-12.57 49.6,-35.01 84.74,-35.01l1460.39 -0.17c35.14,0 68.39,22.61 85.84,35.18l200.47 144.61 0 0zm-12.83 130.44c-15.3,110.89 -115.85,120.47 -230.86,120.47 -115.05,0 -205.22,-9.58 -220.48,-120.47l451.34 0zm-2006.26 0l451.34 0c-15.3,110.89 -110.64,120.47 -225.65,120.47 -115.01,0 -210.39,-9.58 -225.69,-120.47zm1487.91 0c-15.3,110.89 -109.75,120.47 -224.77,120.47 -115.05,0 -211.23,-9.58 -226.53,-120.47l451.3 0zm-518.26 0c-15.3,110.89 -112.36,120.47 -227.38,120.47 -115.05,0 -208.66,-9.58 -223.96,-120.47l451.34 0zm788.52 1141.28l-292.07 -0.17 0.84 -741.26 289.5 -2.53 1.73 743.96zm116.52 -0.08l-71.92 0 0 -770.99c0,-17.28 -6.65,-27.15 -23.92,-27.15l-341 0c-17.28,0 -22.41,9.41 -22.41,26.73l-0.42 771.07 -1283.45 0.34 0 -947.08c39.01,0 61.16,0 94.12,0 113.03,0 211.15,0.13 259.15,-94.37 75.58,82.31 147.68,94.37 260.71,94.37 113.04,0 183.41,-6.05 257.56,-94.37 79.99,81.43 144.44,94.37 257.43,94.37 113.03,0 187.65,-5.17 260.88,-94.37 92.1,78.02 143.55,94.37 256.54,94.37 32.96,0 53.39,0 95.01,0l1.72 947.08 0 0zm-580.94 -792.18c-2.57,-29.59 -11.39,-34.17 -48.64,-33.12l-947.45 0c-37.25,-1.05 -46.71,4.29 -49.27,33.88l0 654.46c2.56,29.64 11.94,34.26 49.18,33.21l947.42 0c37.28,1.05 46.2,-4.58 48.76,-34.21l0 -654.22 0 0zm54.23 -49.1c-2.82,-34.04 -12.53,-39.3 -53.68,-38.12l-1045.78 0c-41.11,-1.18 -51.54,4.92 -54.36,38.97l0 752.61c2.82,34.05 13.12,39.35 54.27,38.17l1045.78 0c41.11,1.18 50.95,-5.29 53.77,-39.34l0 -752.29z"/>
                <g id="_2368313925168">
                </g>
                <g id="_2368313922416">
                </g>
                <g id="_2368313923504">
                </g>
                <g id="_2368313922480">
                </g>
                <g id="_2368313921712">
                </g>
                <g id="_2368313922288">
                </g>
                <g id="_2368313922544">
                </g>
                <g id="_2368313920528">
                </g>
                <g id="_2368313921072">
                </g>
                <g id="_2368313919760">
                </g>
                <g id="_2368313919728">
                </g>
                <g id="_2368313919664">
                </g>
                <g id="_2368313921392">
                </g>
                <g id="_2368313917744">
                </g>
                <g id="_2368313919216">
                </g>
                <path class="fil1 str1" d="M511.26 1532.54c-9.24,19.28 -20.94,34.06 -35.09,44.28 -14.21,10.21 -30.13,15.29 -47.82,15.29 -21.63,0 -39.21,-7.36 -52.67,-22.14 -13.53,-14.78 -20.26,-34.07 -20.26,-57.92 0,-18.2 3.14,-35.09 9.47,-50.73 6.28,-15.69 15.47,-29.22 27.51,-40.74 9.13,-9.02 19.34,-15.81 30.75,-20.43 11.36,-4.62 23.46,-6.96 36.24,-6.96 17.86,0 32.41,4.73 43.77,14.15 11.3,9.41 16.94,21.34 16.94,35.83 0,9.19 -2.34,16.55 -7.07,22.09 -4.74,5.59 -10.96,8.39 -18.66,8.39 -6.56,0 -12.04,-1.94 -16.38,-5.77 -4.33,-3.82 -6.5,-8.67 -6.5,-14.43 0,-3.14 0.85,-6.22 2.51,-9.31 1.65,-3.02 3.99,-5.76 7.02,-8.1 1.42,-1.08 3.76,-2.39 6.9,-3.94 3.2,-1.48 4.79,-2.91 4.79,-4.22 0,-4.05 -3.02,-7.82 -9.01,-11.47 -6.05,-3.59 -13.07,-5.42 -21.12,-5.42 -20.82,0 -37.66,10.16 -50.55,30.42 -12.84,20.25 -19.29,47.02 -19.29,80.23 0,16.49 3.6,29.21 10.78,38.17 7.19,8.96 17.41,13.47 30.65,13.47 10.27,0 20.14,-3.43 29.67,-10.22 9.53,-6.84 18.72,-17 27.5,-30.52l19.92 0zm213.93 -60.26c0,15.12 -2.45,29.9 -7.36,44.28 -4.91,14.38 -11.7,26.99 -20.37,37.95 -9.82,11.92 -21.29,21.11 -34.47,27.56 -13.12,6.44 -26.99,9.7 -41.6,9.7 -21.74,0 -39.31,-7.31 -52.72,-21.97 -13.41,-14.67 -20.15,-33.9 -20.15,-57.75 0,-18.2 3.14,-35.09 9.42,-50.73 6.33,-15.69 15.46,-29.22 27.5,-40.74 9.36,-9.13 19.75,-15.98 31.22,-20.55 11.41,-4.56 23.62,-6.84 36.63,-6.84 21.74,0 39.15,7.19 52.27,21.57 13.07,14.38 19.63,33.55 19.63,57.52zm-100.2 101c8.73,0 16.83,-2.74 24.36,-8.22 7.48,-5.42 14.33,-13.64 20.49,-24.59 5.88,-10.56 10.5,-22.66 13.81,-36.29 3.25,-13.59 4.91,-27.34 4.91,-41.15 0,-17.17 -3.14,-30.01 -9.48,-38.57 -6.27,-8.51 -15.69,-12.78 -28.19,-12.78 -17.8,0 -33.04,10.1 -45.7,30.3 -12.67,20.14 -19.01,45.3 -19.01,75.43 0,17.92 3.37,31.73 10.1,41.37 6.79,9.65 16.32,14.5 28.71,14.5zm175.18 -156.13l-35.61 0 5.2 -18.14 35.2 0c6.8,-24.2 16.1,-42.29 27.97,-54.27 11.81,-11.98 26.13,-17.98 42.96,-17.98 11.02,0 20.03,2.8 27.05,8.34 7.02,5.59 10.56,12.72 10.56,21.4 0,5.82 -1.71,10.49 -5.08,14.15 -3.42,3.59 -7.82,5.42 -13.18,5.42 -5.88,0 -10.56,-1.37 -13.98,-4.17 -3.37,-2.79 -5.08,-6.56 -5.08,-11.35 0,-3.31 1.14,-6.74 3.37,-10.27 0.63,-1.03 0.97,-1.66 0.97,-1.89 0,-0.74 -0.75,-1.48 -2.23,-2.11 -1.54,-0.68 -3.25,-1.03 -5.31,-1.03 -7.98,0 -14.83,4.28 -20.54,12.84 -5.7,8.56 -10.67,22.2 -15.06,40.92l33.61 0 -4.17 18.14 -33.44 0 -32.24 135.24c-8.39,35.1 -18.37,60.32 -30.01,75.61 -11.64,15.35 -26.42,23 -44.23,23 -10.61,0 -19.4,-2.68 -26.42,-8.1 -7.02,-5.43 -10.55,-11.93 -10.55,-19.63 0,-6.57 1.59,-11.81 4.84,-15.7 3.26,-3.93 7.65,-5.87 13.24,-5.87 6,0 10.73,1.42 14.16,4.22 3.36,2.85 5.07,6.79 5.07,11.75 0,3.89 -1.48,7.65 -4.5,11.19 -0.69,0.85 -1.03,1.37 -1.03,1.48 0,0.46 0.91,0.92 2.8,1.43 1.82,0.51 3.93,0.74 6.39,0.74 7.81,0 14.66,-5.99 20.48,-17.92 5.88,-11.98 12.44,-34.06 19.63,-66.25l29.16 -131.19zm120.18 0l-35.61 0 5.19 -18.14 35.21 0c6.79,-24.2 16.09,-42.29 27.96,-54.27 11.81,-11.98 26.14,-17.98 42.97,-17.98 11.01,0 20.03,2.8 27.05,8.34 7.02,5.59 10.55,12.72 10.55,21.4 0,5.82 -1.71,10.49 -5.07,14.15 -3.43,3.59 -7.82,5.42 -13.19,5.42 -5.87,0 -10.55,-1.37 -13.98,-4.17 -3.36,-2.79 -5.08,-6.56 -5.08,-11.35 0,-3.31 1.15,-6.74 3.37,-10.27 0.63,-1.03 0.97,-1.66 0.97,-1.89 0,-0.74 -0.74,-1.48 -2.22,-2.11 -1.55,-0.68 -3.26,-1.03 -5.31,-1.03 -7.99,0 -14.84,4.28 -20.54,12.84 -5.71,8.56 -10.68,22.2 -15.07,40.92l33.61 0 -4.16 18.14 -33.44 0 -32.24 135.24c-8.39,35.1 -18.38,60.32 -30.02,75.61 -11.64,15.35 -26.42,23 -44.22,23 -10.62,0 -19.41,-2.68 -26.42,-8.1 -7.02,-5.43 -10.56,-11.93 -10.56,-19.63 0,-6.57 1.6,-11.81 4.85,-15.7 3.25,-3.93 7.65,-5.87 13.24,-5.87 5.99,0 10.73,1.42 14.15,4.22 3.37,2.85 5.08,6.79 5.08,11.75 0,3.89 -1.48,7.65 -4.51,11.19 -0.68,0.85 -1.03,1.37 -1.03,1.48 0,0.46 0.92,0.92 2.8,1.43 1.83,0.51 3.94,0.74 6.39,0.74 7.82,0 14.67,-5.99 20.49,-17.92 5.87,-11.98 12.44,-34.06 19.63,-66.25l29.16 -131.19zm230.3 110.25c-11.24,21.28 -24.13,37.26 -38.69,47.99 -14.49,10.67 -30.58,16.03 -48.16,16.03 -21.17,0 -38.17,-6.84 -50.95,-20.48 -12.84,-13.64 -19.23,-31.84 -19.23,-54.55 0,-18.61 3.59,-36.01 10.78,-52.27 7.19,-16.21 17.63,-30.59 31.33,-43.2 9.93,-9.02 20.65,-15.98 32.24,-20.83 11.58,-4.85 23.22,-7.25 34.92,-7.25 15.24,0 27.45,3.77 36.58,11.3 9.13,7.54 13.69,17.58 13.69,30.19 0,19.06 -10.55,34.58 -31.67,46.56 -21.11,12.04 -50.9,19.29 -89.47,21.86 -1.31,5.02 -2.28,9.41 -2.91,13.29 -0.63,3.83 -0.92,7.19 -0.92,10.1 0,13.47 3.6,24.2 10.85,32.3 7.24,8.05 16.89,12.1 28.93,12.1 11.47,0 22.54,-3.71 33.27,-11.18 10.67,-7.48 20.6,-18.38 29.73,-32.64l19.68 0.68zm-115.61 -41.43c26.31,-1.65 46.85,-6.9 61.69,-15.75 14.78,-8.78 22.2,-20.03 22.2,-33.61 0,-7.59 -2.12,-13.58 -6.28,-17.97 -4.17,-4.4 -9.82,-6.62 -16.95,-6.62 -12.38,0 -23.85,6.45 -34.41,19.34 -10.61,12.84 -19.34,31.1 -26.25,54.61zm301.87 41.43c-11.24,21.28 -24.14,37.26 -38.69,47.99 -14.5,10.67 -30.59,16.03 -48.16,16.03 -21.17,0 -38.18,-6.84 -50.96,-20.48 -12.84,-13.64 -19.23,-31.84 -19.23,-54.55 0,-18.61 3.59,-36.01 10.78,-52.27 7.19,-16.21 17.64,-30.59 31.33,-43.2 9.93,-9.02 20.66,-15.98 32.24,-20.83 11.59,-4.85 23.23,-7.25 34.93,-7.25 15.23,0 27.44,3.77 36.57,11.3 9.13,7.54 13.7,17.58 13.7,30.19 0,19.06 -10.56,34.58 -31.67,46.56 -21.12,12.04 -50.9,19.29 -89.48,21.86 -1.31,5.02 -2.28,9.41 -2.91,13.29 -0.63,3.83 -0.91,7.19 -0.91,10.1 0,13.47 3.6,24.2 10.84,32.3 7.25,8.05 16.89,12.1 28.93,12.1 11.47,0 22.54,-3.71 33.27,-11.18 10.67,-7.48 20.6,-18.38 29.73,-32.64l19.69 0.68zm-115.61 -41.43c26.3,-1.65 46.85,-6.9 61.68,-15.75 14.78,-8.78 22.2,-20.03 22.2,-33.61 0,-7.59 -2.11,-13.58 -6.28,-17.97 -4.16,-4.4 -9.81,-6.62 -16.94,-6.62 -12.39,0 -23.86,6.45 -34.41,19.34 -10.62,12.84 -19.35,31.1 -26.25,54.61z"/>
                </g>
              </svg>
              h2.thumb-text.text-xs-center ДЕКОРИРОВАНИЕ ВИТРИН
      br.hide-mobile
      br.hide-mobile
      v-layout(row, wrap).subthumb
        v-flex(xs12)
          h2.text-xs-center.header-thumbs Наши преимущества
        v-flex(xs4)
          .thumb
            img(src="@/assets/startup.svg" width="64").svg-icon
            h2.thumb-text.text-xs-center Скорость
        v-flex(xs4)
          .thumb
            img(src="@/assets/handshake.svg" width="64").svg-icon
            h2.thumb-text.text-xs-center Индивидуальный подход
        v-flex(xs4)
          .thumb
            img(src="@/assets/qual.svg" width="64").svg-icon
            h2.thumb-text.text-xs-center Качество
</template>

<script>
import brand from '@/assets/brand.svg'
import decor from '@/assets/decor.svg'
import vinyl from '@/assets/vinyl.svg'

export default {
  data: () => {
    return {
      icons: {
        brand: brand,
        decor: decor,
        vinyl: vinyl
      },
      hover: {
        brand: false,
        decor: false,
        vinyl: false
      },
      thumbLinks: {
        brand: '',
        decor: '',
        vinyl: ''
      }
    }
  },
  methods: {
    mouseOver () {
      this.hover.brand = true
    }
  },
  mounted () {
    if (R.path(['length'], this.links) >= 3) {
      console.table(this.links)
      let [brand, vinyl, decor] = this.links
      this.$set(this, 'thumbLinks', {brand, decor, vinyl})
    }
  },
  props: {
    links: {
      type: Array,
      default: () => {
        return []
      }
    }
  }
}
</script>

<style lang="scss">
  .thumb {
    cursor: pointer;
  }
  $icon-brown: rgb(122, 72, 54);
  .hovering {
    fill: $icon-brown !important;
    color: $icon-brown !important;
  }
  .svg-icon {
    text-align: center;
    display: block;
    margin: 0 auto;
  }
  .thumb-icon {
    display: block;
    margin: 0 auto;
  }
  .icon-vinyl {
    .str0 {stroke:black;stroke-width:12;stroke-miterlimit:22.9256}
    .str1 {stroke:black;stroke-width:0.9;stroke-miterlimit:22.9256}
    .fil0 {fill:none}
    .fil1 {fill:black;fill-rule:nonzero}
  }

  .vinyl-hover svg {
    .str0 {stroke: $icon-brown ;stroke-width:12;stroke-miterlimit:22.9256}
    .str1 {stroke: $icon-brown ;stroke-width:0.9;stroke-miterlimit:22.9256}
    .fil0 {fill:none}
    .fil1 {fill:$icon-brown ;fill-rule:nonzero}
  }
  .center {
    position: relative;
    left: 50%;
    -webkit-transform: translateX(-50%);
    -ms-transform: translateX(-50%);
    transform: translateX(-50%);
  }
  .subthumb {
    .header-thumbs {
      font-size: 1.6em;
    }
    .thumb {
      cursor: default
    }
  }
  @media (min-width: 600px) {
    .subthumb {
      padding-left: 80px;
      padding-right: 80px;
    }
  }
  @media (max-width: 600px) {
    .hide-mobile {
      display: none !important;
      visibility: hidden !important;
    }
    .subthumb {
      .header-thumbs {
        font-size: 1.1em;
      }
      .thumb h2 {
        font-size: 0.9em;
      }
    }
    .thumb-text {
      margin-top: 7px;
    }
  }
</style>
