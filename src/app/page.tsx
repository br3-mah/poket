import Link from 'next/link';
import styles from './landing.module.scss';

export default function LandingPage() {
  return (
    <div className={styles.container}>
      <header className={styles.header}>
        <h1>Pocket Ecommerce</h1>
        <nav>
          <Link href="/login" className={styles.navLink}>Login</Link>
          <Link href="/register" className={styles.navLink}>Register</Link>
        </nav>
      </header>
      
      <main className={styles.hero}>
        <h2>Shop the Future</h2>
        <p>Discover amazing products at unbeatable prices</p>
        <button className={styles.cta}>Browse Products</button>
      </main>
    </div>
  );
}