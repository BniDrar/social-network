// {
//       Path:    "/api/comment/add",
//       handler: app.AddComment,
//       Role:    User,
// },
// {
//       Path:    "/api/comment/get",
//       handler: app.GetComments,
//       Role:    User,
// },


const getComments = async(postId) => {
            try {
              const response = await fetch(
                `${process.env.BACKEND_URL}/api/comment/get?id=${postId}`,
                {
                  method: 'GET',
                  headers: {
                    'Content-Type': 'application/json',
                  },
                  credentials: 'include',
                }
              ); 
              const text = await response.text();
              if (!text) {
                console.warn('Empty response from server');
                return [];
              }
              const data = JSON.parse(text);
              return data;
            } catch (error) {
              console.error('Error fetching posts:', error);
              return [];
            }     
}

const createComment= async (Comment)=> {
      try {
        const formData = new FormData();
        formData.append('content', Comment.content);
      if (Comment.image) {
          formData.append('image', Comment.image);
      }
      formData.append('post_id', Comment.postId);
        const response = await fetch(`${process.env.BACKEND_URL}/api/comment/add`, {
          method: 'POST',
          credentials: 'include',
          body: formData,
        });
        const data = await response.json();
        return {status: response.status, data: data};
      } catch (error) {
        console.error('Error creating comment:', error);
        return {status: 500, error: error.message};
      }
    }


    export {
      getComments,
      createComment
    }