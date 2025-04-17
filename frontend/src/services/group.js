// export const validbackendUrl = `http://localhost:8080`
import { validbackendUrl } from '@/utils/ustil'


async function GetGroup(groupId) {
  try {
    const response = await fetch(`${validbackendUrl}/api/group?id=${groupId}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
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
    const response = await fetch(`${validbackendUrl}/api/group/create`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(groupData),
    });
    const data = await response.json();
    return data;

  } catch (error) {
    console.error('Error creating group:', error);
    throw error;
  }
}

async function GetAllGroups(ReqData) {

  try {
    const response = await fetch(`${validbackendUrl}/api/groups`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(ReqData),
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
    const response = await fetch(`${validbackendUrl}/api/user/groups`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
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
    const response = await fetch(`${validbackendUrl}/api/group/members?id=${groupId}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
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
    const response = await fetch(`${validbackendUrl}/api/group/join/request`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
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
    const response = await fetch(`${validbackendUrl}/api/group/join/response`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
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
    const response = await fetch(`${validbackendUrl}/api/group/invite/request`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(ReqData),
    });
    const data = await response.json();
    return data;
  } catch (error) {
    console.error('Error sending group invite:', error);
    throw error;
  }
}
async function InvitationResponse(ReqData) {
  try {
    const response = await fetch(`${validbackendUrl}/api/group/invite/response`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(ReqData),
    });
    const data = await response.json();
    return data;
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
};
