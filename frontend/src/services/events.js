// type Event struct {
// 	ID          int    `json:"id"`
// 	UserID      int    `json:"user_id"`
// 	Title       string `json:"title"`
// 	Description string `json:"description"`
// 	Date        string `json:"date"`
// 	Time        string `json:"time"`
// 	Location    string `json:"location"`
// 	GroupID     int    `json:"group_id"`
// 	CreatedAt   string `json:"created_at"`
// 	Going       int   `json:"going"`
// }

import { GetGroup } from "./group";


// Event service function
const createEvent = async (eventData) => {
      try {
            const response = await fetch(`${process.env.BACKEND_URL}/api/event/create`, {
                  method: 'POST',
                  body: JSON.stringify(eventData),
                  credentials: 'include',
            });

            const data = await response.json();
            return {
                  status: response.status,
                  data: data,
                  error: !response.ok ? data.error : null
            };
      } catch (error) {
            console.error('Error creating event:', error);
            return {
                  status: 500,
                  error: 'Internal server error'
            };
      }
};

const GetEventsByGroupId = async (groupId) => {
      try {
            const response = await fetch(`${process.env.BACKEND_URL}/api/group/events?id=${groupId}`, {
                  method: 'GET',
                  credentials: 'include',
            });
            const data = await response.json();
            return {
                  status: response.status,
                  data: data,
                  error: !response.ok ? data.error : null
            };
      }catch (error) {
            console.error('Error fetching events:', error);
            return {
                  status: 500,
                  error: 'Internal server error'
            };
      }
}

const getEventById = async (eventId) => {
      try {
            const response = await fetch(`${process.env.BACKEND_URL}/api/event/get?id=${eventId}`, {
                  method: 'GET',
                  credentials: 'include',
            });
            const data = await response.json();
            return {
                  status: response.status,
                  data: data,
                  error: !response.ok ? data.error : null
            };
      } catch (error) {
            console.error('Error fetching event:', error);
            return {
                  status: 500,
                  error: 'Internal server error'
            };
      }
}

const voteEvent = async (eventData)=>{
      try {
            const response = await fetch(`${process.env.BACKEND_URL}/api/event/vote`, {
                  method: 'POST',
                  body: JSON.stringify(eventData),
                  credentials: 'include',
            });
            const data = await response.json();
            return {
                  status: response.status,
                  data: data,
                  error: !response.ok ? data.error : null
            };
      } catch (error) {
            console.error('Error voting for event:', error);
            return {
                  status: 500,
                  error: 'Internal server error'
            };
      }
}

export { createEvent, GetEventsByGroupId, getEventById, voteEvent };