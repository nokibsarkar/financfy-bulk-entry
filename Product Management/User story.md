# ✨ User Stories & Features Documentation

---

## 1️⃣ **User Story: Account Creation**

### 🎯 **Objective**  
Enable users to create personal or business accounts in a secure and reliable manner.

---

### 📝 **Login Requirements**

#### 📧 **Valid Email Address**
- The email address must adhere to a standard format:  
  `username@domain.com`.
- The domain must exist and be associated with a functional email service provider.

#### 🔒 **Strong Password**
- Passwords must be robust and meet the following criteria:  
  - A minimum of **8 characters**.  
  - Includes both **uppercase and lowercase letters** (e.g., "Aa").  
  - Contains at least one **numeric digit** (e.g., "2").  
- Passwords should be easy for the user to remember yet challenging for others to guess.

#### 💾 **Remember Password Option**
- Users will have the option to save their login credentials for a specified number of days, improving user convenience.

---

### 🔐 **Login Actions**

#### ✅ **Sign In**
- Users can log in using their **registered email address** and **password**.

#### 🆕 **Sign Up**
- Users without an existing account must complete the following:  
  1. Provide **Name**, **Email**, and **Password**.  
  2. Reconfirm the password for verification.  
  3. Complete a security step by entering a **6-digit OTP (One-Time Password)** sent to their registered email address.  
  4. An incorrect OTP will prevent account creation, with an option to **resend the OTP**.

---

### 🔄 **Post Account Creation**

#### 🔑 **Login with Email & Password**
- Upon successful registration, users can log in using their **email** and **password**.

#### ❓ **Forgot Password**
1. Initiate a reset by providing the **registered email address**.  
2. A **6-digit OTP** will be sent to the email for verification.  
3. After entering the correct OTP, users can reset their password.  
4. An option to **resend the OTP** will be provided if needed.

---

### ✅ **Acceptance Criteria**
- Users must successfully register by entering all required fields and verifying via OTP.  
- OTP delivery to the registered email must occur within **1 minute** of the request.  
- Incorrect OTPs should display an **error message** with a retry option.  
- Login attempts with incorrect credentials must display appropriate error messages.  
- Password reset must allow secure updates post OTP verification.

---

### 🖥️ **UI Workflow**

#### ✍️ **Sign-Up Flow**
1. Navigate to the **Sign-Up Page**.  
2. Input: **Name**, **Email**, **Password**, and **Confirm Password**.  
3. Submit the form.  
4. Redirect to the **OTP Verification Page**.  
   - Prompt: "Enter the 6-digit OTP sent to your email."  
   - Options: **Submit OTP** | **Resend OTP**  
5. Upon successful OTP verification, registration is complete.

#### 🔓 **Login Flow**
1. Navigate to the **Login Page**.  
2. Input **Email** and **Password**.  
   - Options: **Remember Me** | **Forgot Password**  
3. Successful login redirects to the **Dashboard**.

#### 🔄 **Password Reset Flow**
1. Click **Forgot Password** on the **Login Page**.  
2. Input the registered **Email** to request an OTP.  
3. Enter the OTP on the **Verification Page**.  
   - Options: **Submit OTP** | **Resend OTP**  
4. Redirect to the **Password Reset Page** after successful OTP verification.  
5. Create and confirm a **new password**.  
6. Redirect to the **Login Page**.

---

## 2️⃣ **Dashboard**

### 🎯 **Objective**
Provide users with a centralized hub for features like **Transaction Management**.

---

## 3️⃣ **Transactions**

### 🎯 **Objective**
Organize, track, and monitor daily, monthly, and yearly transactions for personal or business purposes, minimizing errors and ensuring accuracy.

---

### 📋 **Transaction Categories**
- **Cash In**  
- **Cash Out**  
- **Bulk Entries**

---

### 💵 **Cash In & Cash Out Functions**
- Fields:  
  1. **Amount**: Specify the transaction value.  
  2. **Category**: Classify the transaction (e.g., income, expense).  
  3. **Date and Time**: Record when the transaction occurred.  
  4. **Payment Method**: Select one of the following:  
     - Cash  
     - Cheque  
     - Online payment  

- **Save transactions** for future reference.

---

### 📊 **Bulk Entries Functions**
- Fields:  
  1. **Date and Time** of transactions.  
  2. **Type** (e.g., Cash In or Cash Out).  
  3. **Category**.  
  4. **Payment Method**.  
  5. **Remarks**.  

- Users can:  
  - Add New Entries  
  - Remove Existing Entries  

---

### ✅ **Acceptance Criteria**
- Single or bulk transactions must be recorded successfully.  
- Transaction entries should display in an **organized table format** on the dashboard.  
- Users should filter transactions by **date**, **category**, and **payment method**.  
- Deleting or editing a transaction entry must trigger a **confirmation prompt**.  
- Bulk entries must support **CSV imports** (if applicable).

---

### 🖥️ **UI Workflow**

#### 📂 **Transaction Dashboard**
1. Navigate to the **Dashboard**.  
2. Select **Transaction Management** from the menu.  
3. Choose a category:  
   - **Cash In**  
   - **Cash Out**  
   - **Bulk Entries**  

#### 💳 **Cash In/Out Entry Workflow**
1. Click **Add New Transaction** under the respective category.  
2. Fill out the form: **Amount**, **Category**, **Date**, and **Payment Method**.  
3. Click **Save Transaction** to log the entry.

#### 📑 **Bulk Entry Workflow**
1. Click **Bulk Entries** from the Transaction menu.  
2. Use the input form for multiple entries: **Date**, **Type**, **Category**, **Payment Method**, and **Remarks**.  
3. Click **Save** to submit bulk data.

#### ❌ **Delete Workflow**
1. Navigate to the **Transaction Table**.  
2. Click the **Delete icon** next to a transaction.  
3. Confirm the action in the **dialog box**.

---

## 📷 **Reference Images**
- Login  
- Sign-Up  
- Forgot Password  
- Dashboard  
  - Cash In  
  - Cash Out  
  - Bulk Entries  
