#ifndef FUNC_HPP
#define FUNC_HPP

#include "MainWindow.hpp"

#define OK_ERR(condition) (condition ? " OK.\n" : " Error.\n");

QString OkErr(bool condition);
bool LoadDLL();
bool InitGoDLL(QWidget *parent);
void NotifyLoginResult(bool loginResult);

#endif // FUNC_HPP
