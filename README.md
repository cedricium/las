# 📚 Libraries Are Sacred (LAS)

Welcome to Libraries Are Sacred, a gateway to streamlined library management
with a purposefully simple full-stack web app. The name "Libraries Are Sacred"
serves as a clever mnemonic for the app, simplifying the understanding of its
core purpose—a library automation system (LAS) designed to make administrative
tasks effortless.

LAS is designed with a laser focus on simplicity, and this extends to user access.
It's an admin-only application, ensuring that only authorized administrators can
harness its powerful features. However, rest assured, the library catalog remains
open and accessible to the public.

Under the hood, LAS leverages a robust tech stack. The API is powered by Go,
utilizing Gin and GORM to provide efficient backend operations, while the frontend
client is crafted in React, ensuring a smooth and interactive user experience.

**Features:**

- manage administrator access
- add & manage library patrons
- populate book inventory using a dead-simple data import system
- streamline inventory checkouts and returns
- configure checkout restrictions with sane defaults baked in:
  - patron balance exceeding $10
  - 3 outstanding transactions
