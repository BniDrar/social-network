'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import Link from 'next/link';
import styles from './register.module.css';
import { register } from '@/services/auth';

export default function Register() {
  const [formData, setFormData] = useState({
    email: '',
    password: '',
    confirmPassword: '',
    first: '',
    last: '',
    date_of_birth: '',
    nickname: '',
    about_me: '',
    avatar: null,
  });

  const [avatarPreview, setAvatarPreview] = useState('');
  const router = useRouter();

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const handleAvatarChange = (e) => {
    const file = e.target.files[0];
    if (file) {
      setFormData((prev) => ({
        ...prev,
        avatar: file,
      }));
      const reader = new FileReader();
      reader.onloadend = () => {
        setAvatarPreview(reader.result);
      };
      reader.readAsDataURL(file);
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    // Here you would typically send the data to your backend
    const formDataToSend = {
      ...formData,
      status: 0, // private by default
      avatar: formData.avatar,
    }
    let resp = await register(formDataToSend);
    if (resp.status === 201) {
      router.push('/login');
    } else {
      let errorForm = document.querySelector('#errorForm');
      console.log(resp);
      errorForm.innerHTML = resp.error || 'An error occurred';
    }
  };

  return (
    <div className={styles.container}>
      <div className={styles.formWrapper}>
        <h1 className={styles.title}>Create Account</h1>
        <p className={styles.subtitle}>Join us today! Please enter your details</p>

        <form onSubmit={handleSubmit} className={styles.form}>
          <div className={styles.avatarSection}>
            <div className={styles.avatarPreview}>
              {avatarPreview ? (
                <img src={avatarPreview} alt='Avatar preview' className={styles.avatarImage} />
              ) : (
                <div className={styles.avatarPlaceholder}>
                  <span>Upload Photo</span>
                </div>
              )}
            </div>
            <input
              type='file'
              id='avatar'
              name='avatar'
              accept='image/*'
              onChange={handleAvatarChange}
              className={styles.avatarInput}
            />
            <label htmlFor='avatar' className={styles.avatarLabel}>
              Choose Avatar (Optional)
            </label>
          </div>

          <div className={styles.nameGroup}>
            <div className={styles.inputGroup}>
              <input
                type='text'
                name='first'
                placeholder='First Name'
                value={formData.first}
                onChange={handleChange}
                className={styles.input}
                required
              />
            </div>

            <div className={styles.inputGroup}>
              <input
                type='text'
                name='last'
                placeholder='Last Name'
                value={formData.last}
                onChange={handleChange}
                className={styles.input}
                required
              />
            </div>
          </div>

          <div className={styles.inputGroup}>
            <input
              type='email'
              name='email'
              placeholder='Email'
              value={formData.email}
              onChange={handleChange}
              className={styles.input}
              required
            />
          </div>
          <div className={styles.passwordGroup}>
            <div className={styles.inputGroup}>
              <input
                type='password'
                name='password'
                placeholder='Password'
                value={formData.password}
                onChange={handleChange}
                className={styles.input}
                required
              />
            </div>

            <div className={styles.inputGroup}>
              <input
                type='password'
                name='confirmPassword'
                placeholder='Confirm Password'
                value={formData.confirmPassword}
                onChange={handleChange}
                className={styles.input}
                required
              />
            </div>
          </div>

          <div className={styles.inputGroup}>
            <input
              type='date'
              name='date_of_birth'
              value={formData.date_of_birth}
              onChange={handleChange}
              className={styles.input}
              required
            />
          </div>

          <div className={styles.inputGroup}>
            <input
              type='text'
              name='nickname'
              placeholder='Nickname (Optional)'
              value={formData.nickname}
              onChange={handleChange}
              className={styles.input}
            />
          </div>

          <div className={styles.inputGroup}>
            <textarea
              name='about_me'
              placeholder='About Me (Optional)'
              value={formData.about_me}
              onChange={handleChange}
              className={`${styles.input} ${styles.textarea}`}
              rows={4}
            />
          </div>

          <div className={styles.terms}>
            <label>
              <input type='checkbox' required /> I agree to the{' '}
              <Link href='/terms' className={styles.link}>
                Terms of Service
              </Link>{' '}
              and{' '}
              <Link href='/privacy' className={styles.link}>
                Privacy Policy
              </Link>
            </label>
          </div>
          <div id='errorForm' className={styles.error}></div>

          <button type='submit' className={styles.button}>
            Create Account
          </button>
        </form>

        <p className={styles.login}>
          Already have an account?{' '}
          <Link href='/login' className={styles.link}>
            Sign in
          </Link>
        </p>
      </div>
    </div>
  );
}

//  const userData = {
//    nickname: form.nickname.value,
//    email: form.email.value,
//    password: form.password.value,
//    first: form.first.value,
//    last: form.last.value,
//    date_of_birth: form.date_of_birth.value,
//    status: 0, // private by default 1 public
//  };
