async function getPosts(LastPost,lien) {
  try {
    const response = await fetch(
      lien,
      {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body:JSON.stringify(LastPost),
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




async function getPostById(postId) {
  try {
    const response = await fetch(`${process.env.BACKEND_URL}/api/post?post_id=${postId}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include', 
    });

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const contentType = response.headers.get('content-type');
    if (!contentType || !contentType.includes('application/json')) {
      throw new Error("Expected JSON but received something else");
    }

    const data = await response.json();
    return await data;
  } catch (error) {
    console.error('Error fetching post:', error);
    throw error;
  }
}


async function createPost(postData) {
  try {
    const formData = new FormData();
    formData.append('nickname', postData.nickname);
    formData.append('content', postData.content);
    if (postData.image) {
      formData.append('image', postData.image);
    }
    formData.append('status', postData.status);
    if (postData.group_id) {
      formData.append('group_id', postData.group_id);
    }
    if (postData.allowed_viewers) {
      postData.allowed_viewers.forEach((viewerId) => {
        formData.append('allowed_viewers', viewerId);
      });
    }
    const response = await fetch(`${process.env.BACKEND_URL}/api/post/create`, {
      method: 'POST',
      credentials: 'include',
      body: formData,
    });
    const data = await response.json();
    return {status: response.status, data: data};
  } catch (error) {
    console.error('Error creating post:', error);
    return {status: 500, error: error.message};
  }
}

async function likePost(postId){
  try {
    const res = await fetch(`${process.env.BACKEND_URL}/api/post/vote?post_id=${postId}`, {
      method: 'POST',
      credentials: 'include',
    });
    return res
  } catch (error) {
    console.error("Error in likePost:", error);
    throw error;
  }
};

async function GetPostsByGroupId(body) {
  body.limit = body.limit || 10;
  try {
    const response = await fetch(`${process.env.BACKEND_URL}/api/group/posts`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
    
      body: JSON.stringify(body),

      credentials: 'include', 
    });

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const contentType = response.headers.get('content-type');
    if (!contentType || !contentType.includes('application/json')) {
      throw new Error("Expected JSON but received something else");
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error('Error fetching post:', error);
    throw error;
  }
}
  



export {
  getPosts,
  getPostById,
  createPost,
  likePost,
  GetPostsByGroupId
};