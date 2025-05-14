function isValidUrl(url) {
      try {
        if (!url.includes("/media/")){
          return false
        }
        new URL(`${process.env.MEDIA_URL}${url}`);
        return true;
      } catch (error) {
        return false;
      }
    }
export default isValidUrl;