const formatDate = (dateString) => {
      try {
        const date = new Date(dateString);
        return date.toLocaleString();
      } catch (error) {
        console.error("Error formatting date:", error);
        return "Invalid Date";
      }
    };
    
export { formatDate}