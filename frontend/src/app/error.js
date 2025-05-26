'use client'
function Error({ statusCode }) {
  return (
    <div>
      <h1>{statusCode ? `An error ${statusCode} occurred` : 'An error occurred'}</h1>
    </div>
  );
}

Error.getInitialProps = ({ res, err }) => {
  const statusCode = res?.statusCode || err?.statusCode || 500;
  return { statusCode };
};

export default Error;
