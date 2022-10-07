import { uuid } from "uuidv4";
import { createExpense, getExpense, getExpenseLogs, getExpenses, updateExpense } from "./services/expenses/expenses";
import { ExpenseCategory, ExpenseDTO, ReviewStatus } from "./types";

export const CURRENT_TEST_ID = uuid();

async function runTestSuite() {
	// We create the expense

	console.log("Creating expense");
	const createdExpense = await createExpense({
		totalAmount: 10000,
		expenseAt: new Date().toISOString(),
		categorization: ExpenseCategory.SERVICE,
		programId: uuid(),
	});
	console.log("Created expense");
	console.log("\n\n");
	console.log(".......");
	console.log(JSON.stringify(createdExpense));
	console.log(".......");
	console.log("\n\n");

	// We get all expenses
	console.log("Getting all expenses");
	let expenses = await getExpenses();
	if (expenses.length != 1) {
		console.log("\n\n");
		console.log(".......");
		console.log("ERROR : We should only have 1 expense");
		console.log("\n\n");

		return;
	}
	console.log("Got all expense");
	console.log("\n\n");
	console.log(".......");
	console.log(JSON.stringify(expenses));
	console.log(".......");
	console.log("\n\n");

	// We updated our expense
	let updatedExpense: ExpenseDTO = {
		...createdExpense,
		reviewStatus: ReviewStatus.INFO_REQUIRED,
		totalAmount: 20000,
	};

	// We update the expense
	console.log("API : Updating expense");
	updatedExpense = await updateExpense(createdExpense.id, updatedExpense);
	console.log("API : Updated expense");
	console.log("\n\n");
	console.log(".......");
	console.log(JSON.stringify(updatedExpense));
	console.log(".......");
	console.log("\n\n");

	// We verify that we don't get too many expenses
	console.log("API : Getting all expenses");
	expenses = await getExpenses();
	if (expenses.length != 1) {
		console.log("\n\n");
		console.log(".......");
		console.log("ERROR : We should only have 1 expense");
		console.log("\n\n");

		return;
	}
	console.log("API : Got all expense");
	console.log("\n\n");
	console.log(".......");
	console.log(JSON.stringify(expenses));
	console.log(".......");
	console.log("\n\n");

	if (expenses.length != 1) {
		console.log("\n\n");
		console.log(".......");
		console.log("ERROR : We should only have 1 expense");
		console.log("\n\n");

		return;
	}
	if (expenses[0].reviewStatus != ReviewStatus.INFO_REQUIRED) {
		console.log("\n\n");
		console.log(".......");
		console.log("ERROR : Our only expense doens't have the correct review status after modification");
		console.log("\n\n");

		return;
	}
	if (expenses[0].totalAmount != 20000) {
		console.log("\n\n");
		console.log(".......");
		console.log("ERROR : Our only expense doens't have the correct review status after modification");
		console.log("\n\n");

		return;
	}

	// We get our current Expense
	console.log("API : Getting current expense by ID");
	const currentExpense = await getExpense(updatedExpense.id);
	console.log("API : Got current expense");
	console.log("\n\n");
	console.log(".......");
	console.log(JSON.stringify(currentExpense));
	console.log(".......");
	console.log("\n\n");

	// We retrieve all logs for the current expense (We should have 2, for the fields ReviewStatus, TotalAmount)
	console.log("API : Getting current expense logs");
	const currentExpenseLogs = await getExpenseLogs(updatedExpense.id);
	console.log("API : Got current expense logs");

	if (currentExpenseLogs.length != 2) {
		console.log("\n\n");
		console.log(".......");
		console.log("ERROR : We should only have 2 expense logs");
		console.log("\n\n");

		return;
	}

	if (!currentExpenseLogs.find((elt) => elt.field == "ReviewStatus")) {
		console.log("\n\n");
		console.log(".......");
		console.log("ERROR : We should have a log for expense review status change");
		console.log("\n\n");

		return;
	}

	if (!currentExpenseLogs.find((elt) => elt.field == "TotalAmount")) {
		console.log("\n\n");
		console.log(".......");
		console.log("ERROR : We should have a log for expense total amount change");
		console.log("\n\n");

		return;
	}

	console.log("\n\n");
	console.log(".......");
	console.log(JSON.stringify(currentExpenseLogs));
	console.log(".......");
	console.log("\n\n");

	console.log("Works");
}

runTestSuite();
