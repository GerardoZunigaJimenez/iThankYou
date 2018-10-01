-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema iThankYou
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema iThankYou
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `iThankYou` DEFAULT CHARACTER SET utf8 ;
USE `iThankYou` ;

-- -----------------------------------------------------
-- Table `iThankYou`.`subCategories`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `iThankYou`.`subCategories` (
  `subcategory_id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(100) NOT NULL,
  PRIMARY KEY (`subcategory_id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `iThankYou`.`categories`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `iThankYou`.`categories` (
  `category_id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(100) NOT NULL,
  PRIMARY KEY (`category_id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `iThankYou`.`comments`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `iThankYou`.`comments` (
  `comment_id` INT NOT NULL AUTO_INCREMENT,
  `comment` VARCHAR(100) NOT NULL,
  `createdAt` DATETIME NOT NULL,
  `commentscol` VARCHAR(100) NOT NULL,
  PRIMARY KEY (`comment_id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `iThankYou`.`users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `iThankYou`.`users` (
  `user_id` INT NOT NULL AUTO_INCREMENT,
  `last_name` VARCHAR(100) NOT NULL,
  `first_name` VARCHAR(100) NOT NULL,
  `email` VARCHAR(45) NOT NULL,
  `createdAt` DATETIME NOT NULL,
  `updatedAt` DATETIME NULL,
  PRIMARY KEY (`user_id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `iThankYou`.`userThanks`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `iThankYou`.`userThanks` (
  `user_thanks_id` INT NOT NULL AUTO_INCREMENT,
  `user_id` INT NOT NULL,
  `subcategory_id` INT NOT NULL,
  `createdAt` DATETIME NOT NULL,
  `category_id` INT NOT NULL,
  `comment_id` INT NOT NULL,
  PRIMARY KEY (`user_thanks_id`),
  INDEX `subcategory_id_idx` (`subcategory_id` ASC),
  INDEX `category_id_idx` (`category_id` ASC),
  INDEX `comment_id_idx` (`comment_id` ASC),
  INDEX `user_id_idx` (`user_id` ASC),
  CONSTRAINT `subcategory_id`
    FOREIGN KEY (`subcategory_id`)
    REFERENCES `iThankYou`.`subCategories` (`subcategory_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `category_id`
    FOREIGN KEY (`category_id`)
    REFERENCES `iThankYou`.`categories` (`category_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `comment_id`
    FOREIGN KEY (`comment_id`)
    REFERENCES `iThankYou`.`comments` (`comment_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `user_id`
    FOREIGN KEY (`user_id`)
    REFERENCES `iThankYou`.`users` (`user_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;

