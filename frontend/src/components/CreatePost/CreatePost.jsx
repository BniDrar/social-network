import styles from "./CreatePost.module.css"
import Form  from "next/form"

/*
    - content
    - image
*/

const CreatePost = () => {
    const HandleForm = async (formData) => {
        "use server"
        const content = formData.get("content")
        const image = formData.get("image")
        console.log(content);
        console.log(image);
    }
    return (
        <>
            <div className={styles.overlay} id="overlay"></div>
            <div className={styles.card}>
                <h3 className={styles.cardTitle}>Create Post</h3>
                <div className={styles.cardBody}>
                    <Form action={HandleForm}>
                        <textarea name="content" placeholder="What's on your mind..." className={styles.formContent}></textarea>
                        <div className={styles.formGroup}>
                            <select name="status" className={styles.status}>
                                <option value="public">Public</option>
                                <option value="private">Private</option>
                                <option value="almost-private">Almost Private</option>
                            </select>
                            <label htmlFor="imageUpload" className={styles.imgUploadLabel}>Upload 📷</label>
                            <input type="file" name="image" className={styles.formImage} id="imageUpload" />
                        </div>
                        <hr className={styles.line} />
                        <button type="submit" className={styles.formSubmit}>submit</button>
                    </Form>
                </div>
                <div className={styles.cardFooter}></div>
            </div>
        </>
    );
}

export default CreatePost;