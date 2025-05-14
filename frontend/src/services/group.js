async function GetGroup(groupId) {
  try {
    const response = await fetch(`${process.env.BACKEND_URL}/api/group?id=${groupId}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include',
    });
    const data = await response.json();
    return data;

  } catch (error) {
    console.error('Error fetching group:', error);
    throw error;
  }
}

async function CreateGroup(groupData) {
  try {
    const response = await fetch(`${process.env.BACKEND_URL}/api/group/create`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include',
      body: JSON.stringify(groupData),
    });
    const data = await response.json();
    if (!response.ok) {
      const error = (data && data.error) || response.statusText;
      return { status: response.status, error: error };
    }

    return { status: response.status, data: data };

  } catch (error) {
    console.error('Error creating group:', error);
    return { status: 500, error: 'An unexpected error occurred' };
  }
}

async function GetAllGroups() {

  try {
    const response = await fetch(`${process.env.BACKEND_URL}/api/groups`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include',
    });
    const data = await response.json();
    return data;
  } catch (error) {
    console.error('Error fetching groups:', error);
    throw error;
  }
}

async function GetUserGroups() {
  try {
    const response = await fetch(`${process.env.BACKEND_URL}/api/user/groups`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include',
    });
    const data = await response.json();
    return data;
  } catch (error) {
    console.error('Error fetching groups:', error); //
    throw error;
  }
}

async function GetGroupMembers(groupId) {
  try {
    const response = await fetch(`${process.env.BACKEND_URL}/api/group/members?id=${groupId}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include',
    });
    const data = await response.json();
    return data;
  } catch (error) {
    console.error('Error fetching groups:', error); //
    throw error;
  }
}

async function RequestToJoinGroup(groupId) {
  try {
    const response = await fetch(`${process.env.BACKEND_URL}/api/group/join/request`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include',
      body: JSON.stringify({ groupId }),
    });
    const data = await response.json();
    return data;
  } catch (error) {
    console.error('Error sending request to join group:', error);
    throw error;
  }
}

async function RequestToJoinResponse(data) {
  try {
    const response = await fetch(`${process.env.BACKEND_URL}/api/group/join/response`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include',
      body: JSON.stringify(data),
    });
    const data = await response.json();
    return data;
  } catch (error) {
    console.error('Error accepting group join request:', error);
    throw error;
  }
}


async function InviteToJoinGroup(ReqData) {
  try {
    const response = await fetch(`${process.env.BACKEND_URL}/api/group/invite/request`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include',
      body: JSON.stringify(ReqData),
    });
    return response;
  } catch (error) {
    console.error('Error sending group invite:', error);
    throw error;
  }
}
async function InvitationResponse(ReqData) {
  try {
    const response = await fetch(`${process.env.BACKEND_URL}/api/group/invite/response`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include',
      body: JSON.stringify(ReqData),
    });
    return response
    } catch (error) {
    console.error('Error accepting group invite:', error);
    throw error;
  }
}

async function GetSuggestedUsers(groupId) {
  try {
    const response = await fetch(`${process.env.BACKEND_URL}/api/group/suggested_users?id=${groupId}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include',
    });
    const data = await response.json();
    return data;
  } catch (error) {
    console.error('Error fetching groups:', error); 
    throw error;
  }
}

const acceptFollowRequest = async (ReqData) => {
  try {
    const response = await fetch(`${process.env.BACKEND_URL}/api/user/follow/response`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include',
      body: JSON.stringify(ReqData),
    });
    return response
  } catch (error) {
    console.error('Error accepting group invite:', error);
    throw error;
  }
}

export {
  GetGroup,
  CreateGroup,
  GetAllGroups,
  GetGroupMembers,
  GetUserGroups,
  RequestToJoinGroup,
  RequestToJoinResponse,
  InviteToJoinGroup,
  InvitationResponse,
  GetSuggestedUsers,
  acceptFollowRequest,
};
