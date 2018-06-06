import styled, { injectGlobal } from 'styled-components';

export const Root = styled.div`
  text-align: center;
`;

export const PageWrapper = styled.div`
  font-size: large;
`;

// native fonts
const font = `-apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen, Ubuntu, Cantarell, "Fira Sans", "Droid Sans", "Helvetica Neue", Arial, sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", sans-serif`;

export const setGlobalStyles = () => {
  injectGlobal`
    body {
      font-family: ${font};
    }

    * {
      box-sizing: border-box;
    }

    body {
      margin: 0;
      padding: 0;
    }
  `;
};
