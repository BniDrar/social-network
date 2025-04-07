import styles from "./CreatePost.module.css"
import Form  from "next/form"

/*
    - content
    - image
*/

const CreatePost = () => {
    const HandleForm = async (formData) => {

    }
    return (
        <>
            <div className={styles.overlay} id="overlay"></div>
            <div className={styles.card}>
                <h3 className={styles.cardTitle}>Create Post</h3>
                <div className={styles.cardBody}>
                    <Form action='HandleForm'>
                        <textarea name="content" placeholder="What's on your mind..." className={styles.formContent}></textarea>
                        <label htmlFor="imageUpload" className={styles.imgUploadLabel}>Upload 📷</label>
                        <input type="file" name="image" className={styles.formImage} id="imageUpload" />
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