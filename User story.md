# User Story: Account Creation 🚀

## **Objective**  
Enable users to create personal or business accounts securely and reliably.

---

## **Login Requirements**  
### **Valid Email Address**  
- Must follow the standard format: `username@domain.com`.
- The domain must be valid and associated with a functional email service provider.

### **Strong Password**  
- Minimum of 8 characters.  
- Includes both uppercase and lowercase letters (e.g., `Aa`).  
- Contains at least one numeric digit (e.g., `2`).  
- Easy for the user to remember but hard for others to guess.  

### **Remember Password Option**  
- Users can opt to save their login credentials for convenience over a specified duration.

---

## **Login Actions**  
### **Sign In**  
- Users log in with their registered email and password.  

### **Sign Up**  
- Users without an account must complete the following:  
  1. Provide Name, Email, and Password.  
  2. Reconfirm the Password.  
  3. Complete a security step by entering a 6-digit OTP sent to their registered email.  
     - Incorrect OTP blocks account creation, but users can resend the OTP.  

---

## **Post Account Creation**  
- **Successful Registration:** Users log in with their Email and Password.  
- **Forgot Password:**  
  1. Enter the registered Email.  
  2. Receive a 6-digit OTP for verification.  
  3. Reset Password after entering the correct OTP.  
  4. Option to resend OTP is available.

---

## **Acceptance Criteria**  
- Users must register successfully after entering all required fields and OTP verification.  
- OTP delivery should occur within 1 minute of request.  
- Incorrect OTP prompts error messages with retry options.  
- Login should fail for incorrect Email/Password and display error messages.  
- Password reset must allow secure credential updates post OTP verification.  

---

## **UI Workflow**  

### **Sign-Up Flow**  
1. Navigate to the Sign-Up Page.  
2. Input:  
   - Name  
   - Email  
   - Password  
   - Confirm Password  
3. Submit the form.  
4. Redirect to OTP Verification Page:  
   - Enter the 6-digit OTP.  
   - Options: **Submit OTP** or **Resend OTP**  
5. Successful OTP verification completes registration.

### **Login Flow**  
1. Navigate to the Login Page.  
2. Input Email and Password.  
   - Options:  
     - **Remember password**  
     - **Forgot Password**  
3. Successful login redirects to the Dashboard.  

### **Password Reset Flow**  
1. Click **Forgot Password** on the Login Page.  
2. Input registered Email to request OTP.  
3. Enter OTP on the Verification Page.  
   - Options: **Submit OTP** or **Resend OTP**  
4. Redirect to Password Reset Page.  
5. Set a new Password and Confirm.  
6. Redirect to Login Page.  

---

# Dashboard: Transaction Management 🖥️

## **Objective**  
Enable users to track and monitor transactions effectively to minimize errors and enhance reconciliation.

---

## **Transaction Categories**  
- **Cash In**  
- **Cash Out**  
- **Bulk Entries**

---

## **Cash In & Cash Out Functions**  
### **Fields:**  
1. **Amount** - Specify the transaction value.  
2. **Category** - Classify the transaction (e.g., income, expense).  
3. **Date and Time** - Record transaction timing.  
4. **Payment Method:**  
   - Cash  
   - Cheque  
   - Online payment  

### **Features:**  
- Save transactions for future reference.

---

## **Bulk Entries Functions**  
### **Fields:**  
1. Date and Time  
2. Type (Cash In or Cash Out)  
3. Category  
4. Payment Method  
5. Remarks  

### **Actions:**  
- Add New Entries  
- Remove Existing Entries  

---

## **Acceptance Criteria**  
- Single/Bulk transactions must be recorded successfully.  
- Transaction entries should display in an organized table format on the dashboard.  
- Users can filter transactions by date, category, and payment method.  
- Delete/Edit actions must prompt confirmation.  
- Bulk entries should support CSV imports (if applicable).  

---

## **UI Workflow**  

### **Transaction Dashboard**  
1. Navigate to the Dashboard.  
2. Select **Transaction Management** from the menu.  
3. Choose a category:  
   - Cash In  
   - Cash Out  
   - Bulk Entries  

### **Cash In/Out Workflow**  
1. Click **Add New Transaction** under the respective category.  
2. Fill out the form:  
   - Amount  
   - Category  
   - Date  
   - Payment Method  
3. Click **Save Transaction** to log the entry.

### **Bulk Entries Workflow**  
1. Click **Bulk Entries** from the menu.  
2. Use the form for multiple entries:  
   - Date  
   - Type  
   - Category  
   - Payment Method  
   - Remarks  
3. Click **Save** to submit bulk data.

### **Delete Workflow**  
1. Navigate to the Transaction Table.  
2. Click the **Delete** icon next to a transaction.  
3. Confirm the action in the dialog box.

---


